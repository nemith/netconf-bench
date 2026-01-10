package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	gossh "golang.org/x/crypto/ssh"
	"nemith.io/netconf"
	"nemith.io/netconf/rpc"
	"nemith.io/netconf/transport/ssh"
)

var (
	host   = flag.String("host", "localhost", "NETCONF server host")
	port   = flag.Int("port", 8830, "NETCONF server port")
	user   = flag.String("user", "admin", "SSH username")
	pass   = flag.String("pass", "admin", "SSH password")
	size   = flag.Int("size", 1024, "Response size in bytes")
	count  = flag.Int("count", 10, "Number of requests")
	method = flag.String("method", "exec", "Request method: exec or do")
)

func main() {
	flag.Parse()

	sshConfig := &gossh.ClientConfig{
		User: *user,
		Auth: []gossh.AuthMethod{
			gossh.Password(*pass),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	target := fmt.Sprintf("%s:%d", *host, *port)

	slog.Info("connected", "server", target)

	ctx := context.Background()
	transport, err := ssh.Dial(ctx, "tcp", target, sshConfig)
	if err != nil {
		slog.Error("failed to connect", "error", err)
		os.Exit(1)
	}
	defer transport.Close()

	session, err := netconf.NewSession(transport)
	if err != nil {
		slog.Error("failed to create netconf session", "error", err)
		os.Exit(1)
	}
	defer func() {
		slog.Info("closing session")
		session.Close(ctx)
	}()

	slog.Info("running requests", "count", *count, "size", *size, "method", *method)

	start := time.Now()

	if *method == "do" {
		runSequentialDo(ctx, session)
	} else {
		runSequentialExec(ctx, session)
	}

	duration := time.Since(start)

	fmt.Printf("%.3f\n", duration.Seconds())

	throughput := float64(*count) / duration.Seconds()
	dataTransferred := float64(*count**size) / (1024 * 1024)
	slog.Info("completed", "duration", duration, "rps", throughput, "transfered_MB", dataTransferred)
}

func runSequentialExec(ctx context.Context, session *netconf.Session) {
	for i := range *count {
		logger := slog.With("req_id", i)

		filter := rpc.SubtreeFilter(fmt.Sprintf("<size>%d</size>", *size))
		req := &rpc.Get{Filter: filter}

		logger.Info("sending request", "size", *size)
		var reply rpc.GetReply
		if err := session.Exec(ctx, req, &reply); err != nil {
			logger.Error("request failed", "error", err)
		}
		logger.Info("request completed")
	}
}

func runSequentialDo(ctx context.Context, session *netconf.Session) {
	for i := range *count {
		logger := slog.With("req_id", i)

		filter := rpc.SubtreeFilter(fmt.Sprintf("<size>%d</size>", *size))
		req := &rpc.Get{Filter: filter}
		rpcMsg := netconf.NewRPC(req)

		logger.Info("sending request", "size", *size)
		msg, err := session.Do(ctx, rpcMsg)
		if err != nil {
			logger.Error("request failed", "error", err)
			continue
		}
		msg.Close()
		logger.Info("request completed")
	}
}
