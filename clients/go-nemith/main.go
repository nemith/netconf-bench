package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	gossh "golang.org/x/crypto/ssh"
	"nemith.io/netconf"
	"nemith.io/netconf/rpc"
	"nemith.io/netconf/transport/ssh"
)

func main() {
	var (
		addr        string
		size, count int
		method      string
		framing     string
	)

	flag.StringVar(&addr, "addr", "localhost:8830", "NETCONF server address")
	flag.IntVar(&size, "size", 1024, "Response size in bytes")
	flag.IntVar(&count, "count", 10, "Number of requests")
	flag.StringVar(&method, "method", "do", "do or exec for the method to call")
	flag.StringVar(&framing, "framing", "", "Framing type: marked or chunked")

	flag.Parse()

	switch framing {
	case "marked", "chunked":
	default:
		slog.Error("invalid framing type", "framing", framing)
		os.Exit(1)
	}

	sshConfig := &gossh.ClientConfig{
		User: "user",
		Auth: []gossh.AuthMethod{
			gossh.Password("pass"),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	slog.Info("connected", "server", addr)

	ctx := context.Background()
	transport, err := ssh.Dial(ctx, "tcp", addr, sshConfig)
	if err != nil {
		slog.Error("failed to connect", "error", err)
		os.Exit(1)
	}
	defer transport.Close()

	slog.Info("running requests", "count", count, "size", size, "method", method)

	var fn func(ctx context.Context, session *netconf.Session, size, count int)

	switch method {
	case "do":
		fn = runDo
	case "exec":
		fn = runExec
	default:
		slog.Error("unknown method", "method", method)
		os.Exit(1)
	}

	startSetup := time.Now()
	session, err := netconf.NewSession(transport)
	if err != nil {
		slog.Error("failed to create netconf session", "error", err)
		os.Exit(1)
	}
	defer func() {
		slog.Info("closing session")
		session.Close(ctx)
	}()
	setupDur := time.Since(startSetup)

	start := time.Now()
	fn(ctx, session, size, count)
	duration := time.Since(start)

	// impl,framing,size,count,setup_ms,rpc_calls_ms,rpc
	fmt.Printf("Go: nemith.io/netconf (%s),%s,%d,%d,%d,%d,%.3f\n",
		strings.Title(method),
		framing,
		size,
		count,
		setupDur.Milliseconds(),
		duration.Milliseconds(),
		float64(count)/duration.Seconds())
}

func runExec(ctx context.Context, session *netconf.Session, size, count int) {
	for i := range count {
		logger := slog.With("req_id", i)

		filter := rpc.SubtreeFilter(fmt.Sprintf("<size>%d</size>", size))
		req := &rpc.Get{Filter: filter}

		logger.Info("sending request", "size", size)
		var reply rpc.GetReply
		if err := session.Exec(ctx, req, &reply); err != nil {
			logger.Error("request failed", "error", err)
		}
		logger.Info("request completed")
	}
}

func runDo(ctx context.Context, session *netconf.Session, size, count int) {
	for i := range count {
		logger := slog.With("req_id", i)

		filter := rpc.SubtreeFilter(fmt.Sprintf("<size>%d</size>", size))
		req := &rpc.Get{Filter: filter}
		rpcMsg := netconf.NewRPC(req)

		logger.Info("sending request", "size", size)
		msg, err := session.Do(ctx, rpcMsg)
		if err != nil {
			logger.Error("request failed", "error", err)
			continue
		}
		msg.Close()
		logger.Info("request completed")
	}
}
