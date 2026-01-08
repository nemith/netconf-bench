package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"golang.org/x/crypto/ssh"
)

// --- Constants ---
const (
	MsgTerminator10 = "]]>]]>"

	ServerHello10 = `<?xml version="1.0" encoding="UTF-8"?>
<hello xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <capabilities>
    <capability>urn:ietf:params:netconf:base:1.0</capability>
  </capabilities>
  <session-id>%d</session-id>
</hello>`

	ServerHello11 = `<?xml version="1.0" encoding="UTF-8"?>
<hello xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <capabilities>
    <capability>urn:ietf:params:netconf:base:1.0</capability>
    <capability>urn:ietf:params:netconf:base:1.1</capability>
  </capabilities>
  <session-id>%d</session-id>
</hello>`

	RpcReplyHeader = `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="%s">`
	RpcReplyFooter = `</rpc-reply>`
	OkReply        = `<ok/>`
)

var (
	port       = flag.Int("port", 8830, "SSH port to listen on")
	useChunked = flag.Bool("chunked", true, "Enable NETCONF 1.1 chunked framing")
	sessID     atomic.Int32
)

// --- Buffer Pool ---
var bufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 65536))
	},
}

// --- Robust Framing ---

type NetconfCodec struct {
	r       *bufio.Reader
	w       *bufio.Writer
	chunked bool
}

func NewCodec(r io.Reader, w io.Writer) *NetconfCodec {
	return &NetconfCodec{
		// Single Shared Reader to prevent buffer loss during mode switch
		r: bufio.NewReaderSize(r, 65536),
		w: bufio.NewWriterSize(w, 65536),
	}
}

func (c *NetconfCodec) UpgradeToChunked() {
	c.chunked = true
}

func (c *NetconfCodec) Write(data []byte) error {
	if !c.chunked {
		if _, err := c.w.Write(data); err != nil {
			return err
		}
		if _, err := c.w.WriteString(MsgTerminator10); err != nil {
			return err
		}
	} else {
		// NETCONF 1.1 Chunked: \n#LEN\n data \n##\n
		fmt.Fprintf(c.w, "\n#%d\n", len(data))
		c.w.Write(data)
		c.w.WriteString("\n##\n")
	}
	return c.w.Flush()
}

func (c *NetconfCodec) ReadMsg() ([]byte, error) {
	if !c.chunked {
		// NETCONF 1.0 (Robust Read)
		// We read byte-by-byte or slice-by-slice looking for ]]>]]>
		// We avoid Scanner to ensure we don't over-buffer and lose data on switch.
		var msg []byte
		suffix := []byte(MsgTerminator10)
		for {
			// Read until '>' to speed up search
			chunk, err := c.r.ReadSlice('>')
			if err != nil {
				if err == bufio.ErrBufferFull {
					// Buffer is full, append what we have and keep reading
					msg = append(msg, chunk...)
					continue
				}
				if err == io.EOF && len(msg) > 0 {
					return msg, nil // Return what we have (lax)
				}
				return nil, err
			}
			msg = append(msg, chunk...)

			// Check if we hit the terminator
			if len(msg) >= 6 && bytes.Equal(msg[len(msg)-6:], suffix) {
				return msg[:len(msg)-6], nil
			}
		}
	}

	// NETCONF 1.1 Chunked Reader
	payload := bufferPool.Get().(*bytes.Buffer)
	payload.Reset()
	defer bufferPool.Put(payload)

	for {
		// 1. Skip potential newlines before chunk start
		// (The RFC says fragments are back-to-back, but implementation variance exists)
		// We look specifically for \n#

		// Read until we find the start marker format
		for {
			b, err := c.r.ReadByte()
			if err != nil {
				return nil, err
			}

			// Start of chunk is \n#
			if b == '\n' {
				next, err := c.r.Peek(1)
				if err != nil {
					return nil, err
				}
				if next[0] == '#' {
					c.r.ReadByte() // Consume '#'
					break
				}
			}
			// If it's not the start marker, we technically ignore it or it's framing error.
			// In strict 1.1, bytes between chunks are invalid, but we just consume.
		}

		// 2. Read Length or Hash
		line, err := c.r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		token := strings.TrimSpace(line)

		if token == "#" { // End of chunks (\n##\n)
			// Return copy because buffer goes back to pool
			res := make([]byte, payload.Len())
			copy(res, payload.Bytes())
			return res, nil
		}

		length, _ := strconv.Atoi(token)
		if length > 0 {
			if _, err := io.CopyN(payload, c.r, int64(length)); err != nil {
				return nil, err
			}
		}
	}
}

// --- XML Structures ---

type RPCMessage struct {
	XMLName   xml.Name  `xml:"rpc"`
	MessageID string    `xml:"message-id,attr"`
	Get       *GetBody  `xml:"get"`
	Close     *struct{} `xml:"close-session"`
}

type GetBody struct {
	Filter *FilterBody `xml:"filter"`
}

type FilterBody struct {
	// Captures <size>123</size> inside <filter>
	Size string `xml:"size"`
}

// --- Handler ---

func handleSession(logger *slog.Logger, channel ssh.Channel, id int) {
	defer channel.Close()
	codec := NewCodec(channel, channel)

	// Send Hello (1.0 Framing) with appropriate capabilities
	var serverHello string
	if *useChunked {
		serverHello = ServerHello11
	} else {
		serverHello = ServerHello10
	}
	fmt.Fprintf(channel, serverHello+MsgTerminator10, id)

	// Read Client Hello
	clientHelloBytes, err := codec.ReadMsg()
	if err != nil {
		logger.Error("failed_read_hello", "err", err)
		return
	}

	// Upgrade Check - only upgrade if server supports chunked and client supports 1.1
	if *useChunked && strings.Contains(string(clientHelloBytes), ":base:1.1") {
		codec.UpgradeToChunked()
	}

	for {
		reqBytes, err := codec.ReadMsg()
		if err != nil {
			if err != io.EOF {
				logger.Error("read_error", "err", err)
			}
			return
		}

		if len(reqBytes) == 0 {
			continue
		}

		var rpc RPCMessage
		if err := xml.Unmarshal(reqBytes, &rpc); err != nil {
			logger.Error("xml_error", "err", err)
			continue
		}

		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Reset()

		fmt.Fprintf(buf, RpcReplyHeader, rpc.MessageID)

		if rpc.Close != nil {
			buf.WriteString(OkReply)
			buf.WriteString(RpcReplyFooter)
			codec.Write(buf.Bytes())
			bufferPool.Put(buf)
			return
		}

		if rpc.Get != nil {
			reqSize := 1024
			// Safe integer parsing
			if rpc.Get.Filter != nil && rpc.Get.Filter.Size != "" {
				// Handle potential whitespace inside XML tag
				cleanSize := strings.TrimSpace(rpc.Get.Filter.Size)
				if s, err := strconv.Atoi(cleanSize); err == nil && s > 0 {
					reqSize = s
				}
			}
			slog.Info("sending response", "size", reqSize)
			generateData(buf, reqSize)
		} else {
			buf.WriteString(OkReply)
		}

		buf.WriteString(RpcReplyFooter)

		if err := codec.Write(buf.Bytes()); err != nil {
			bufferPool.Put(buf)
			return
		}

		bufferPool.Put(buf)
	}
}

func generateData(buf *bytes.Buffer, size int) {
	buf.WriteString("<data><interfaces>")
	// Pre-calculated template for speed
	tmpl := "<interface><name>eth%d</name><mtu>1500</mtu><status>up</status></interface>"
	curr := 0
	// 50 bytes for header/footer safety
	limit := size - 50
	for buf.Len() < limit {
		fmt.Fprintf(buf, tmpl, curr)
		curr++
	}
	buf.WriteString("</interfaces></data>")
}

// --- Main ---

func main() {
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	key, err := ssh.ParsePrivateKey([]byte(testHostKey))
	if err != nil {
		panic(err)
	}

	config := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) { return nil, nil },
	}
	config.AddHostKey(key)

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic(err)
	}

	logger.Info("server_started", "port", *port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			_, chans, reqs, err := ssh.NewServerConn(c, config)
			if err != nil {
				return
			}
			go ssh.DiscardRequests(reqs)

			for newChannel := range chans {
				if newChannel.ChannelType() != "session" {
					newChannel.Reject(ssh.UnknownChannelType, "unknown")
					continue
				}
				channel, requests, err := newChannel.Accept()
				if err != nil {
					continue
				}

				// Non-blocking SSH Request Loop
				go func(in <-chan *ssh.Request) {
					for req := range in {
						if req.Type == "subsystem" && string(req.Payload[4:]) == "netconf" {
							req.Reply(true, nil)
							// Spawn session async
							sid := int(sessID.Add(1))
							go handleSession(logger, channel, sid)
						} else {
							req.Reply(false, nil)
						}
					}
				}(requests)
			}
		}(conn)
	}
}

const testHostKey = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn
NhAAAAAwEAAQAAAQEAxCL24d3v+/7HorZZsQZMOaUaZxLuKIAQdfr85oJfEaOQ8kmn2x27
DBv6n93ba4LCrRRQgTr65ZiEvLyEa2eW6z6zCphaK6S3tnUoc3jNpL/pTF32CZTgfVJ1lB
KoSnXGwCwM8upfnEDktjm/H8itIFkOpEtEHOgZp2gKby82y53z7+LUUV3iHcNPPsRRJT/P
QFCT7O2IASmbhLdFCQ28UUQLSdPt9ilFecW4AmXL6ORHt6LQyix0UWky0L0oQUx/lA3dNA
ZleK7w5AtBPivX4cE8zrAKPEo+i0xuALbZ609/l6VrVEZ2aFsuSuYxQqdJQADvmgiPMw3n
DpbPedzZowAAA9CDiiZRg4omUQAAAAdzc2gtcnNhAAABAQDEIvbh3e/7/seitlmxBkw5pR
pnEu4ogBB1+vzmgl8Ro5DySafbHbsMG/qf3dtrgsKtFFCBOvrlmIS8vIRrZ5brPrMKmFor
pLe2dShzeM2kv+lMXfYJlOB9UnWUEqhKdcbALAzy6l+cQOS2Ob8fyK0gWQ6kS0Qc6BmnaA
pvLzbLnfPv4tRRXeIdw08+xFElP89AUJPs7YgBKZuEt0UJDbxRRAtJ0+32KUV5xbgCZcvo
5Ee3otDKLHRRaTLQvShBTH+UDd00BmV4rvDkC0E+K9fhwTzOsAo8Sj6LTG4AttnrT3+XpW
tURnZoWy5K5jFCp0lAAO+aCI8zDecOls953NmjAAAAAwEAAQAAAQEAppztDEH0MyTjgZ1V
48NtzSorm5PBLDZdxVtIqqflCp7f6nIXXVOMKWU7KcLVOicCKPUwzhbEO3WsjIe4FWUSFx
RUE5QIRTrPtPchPbZJOsKr9Gt9LfaCHPOHXhnbTCVwiJ7jAReH0LxvjTzQ1rcqGmiMv5QQ
cF3aLRN8iaaZ7azp9YbTvQHb5nOpkG7k5CTPdZ23pXQVsDPPVRAgeZdqfZHMM0aPG/t1jy
DjvmI4xTVWX2Qw4WtWm66Ag9vW4aeTvrDVTer3FZXIde+RkoLrFpEhMJdJ1yqogpOeqo7S
5v2kfk5QoiX+BIYCFyrtodaHipC9JxPke9KiMvMFscphQQAAAIBGzqz454NkhMCKkrJsV/
3+3XQ+EDP5swEMXUa3B5wTLy57JR2V9CFb9fQTc0RIb221N5qTixEiztxwG+FxlsqcBJl1
RF/9Mmx7xpRsLRxseVo0I37qR5jFjNYfvlPSvSG0MoNjhZn6yBKxBVNNEwQesl2fwLKo0I
Ro1CnF87pbTAAAAIEA9LXfX0mJ1WpHJrN6PqFsltGNFcp6SftDApCbhHkX8KwbtI/ZDl2s
vD4jPXiIOoJS2PAmZmH4M0njVKFX7HkZ/5jS3g8nsGdRv0Z3Oe/ilzuhXFpry15qlao3Ct
MI9byJkxI57ZEFh7YetPHfdlv/mgto+DrqKpUokfATpO+lSaUAAACBAM0vak63wRmmD05O
LygYSgaJ2g5ARGXqyC1GOCBPi/pEm+pBYT2Jpkg39o9v1Glz5MZ+hpq4H2VORwQ/rAkjU1
NnU3p1na6N1BkZLpv9Jn/9N2vnxnXChTPDMY96baLlYXrLcwArjQHG7cGloVwwim62hg1o
Ylvc4QmbV9QII2OnAAAAGGJiZW5uZXR0QENXLUhNOUQ0TVFNVzItTAEC
-----END OPENSSH PRIVATE KEY-----`
