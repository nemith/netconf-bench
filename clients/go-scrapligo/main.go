package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/driver/options"
)

func main() {
	var (
		addr          string
		size, count   int
		transportType string
		framing       string
	)

	flag.StringVar(&addr, "addr", "localhost:8830", "NETCONF server address")
	flag.IntVar(&size, "size", 1024, "Response size in bytes")
	flag.IntVar(&count, "count", 10, "Number of requests")
	flag.StringVar(&transportType, "transport", "standard", "Transport type system (/usr/bin/ssh) or standard (crypto/ssh)")
	flag.StringVar(&framing, "framing", "", "Framing type: marked or chunked")
	flag.Parse()

	switch framing {
	case "marked", "chunked":
	default:
		slog.Error("invalid framing type", "framing", framing)
		os.Exit(1)
	}

	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		slog.Error("failed to parse address", "error", err)
		os.Exit(1)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		slog.Error("failed to parse port", "error", err)
		os.Exit(1)
	}

	if transportType != "standard" && transportType != "system" {
		slog.Error("invalid transport type", "type", transportType)
		os.Exit(1)
	}

	slog.Info("connecting", "server", addr)

	setupStart := time.Now()
	driver, err := netconf.NewDriver(
		host,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("user"),
		options.WithAuthPassword("pass"),
		options.WithTransportType(transportType),
		options.WithPort(port),
	)
	if err != nil {
		slog.Error("failed to create driver", "error", err)
		os.Exit(1)
	}

	if err := driver.Open(); err != nil {
		slog.Error("failed to connect", "error", err)
		os.Exit(1)
	}
	defer func() {
		slog.Info("closing session")
		driver.Close()
	}()
	setupDur := time.Since(setupStart)

	slog.Info("running requests", "count", count, "size", size)

	start := time.Now()
	run(driver, size, count)
	duration := time.Since(start)

	// name,framing,size,count,setup_ms,rpc_calls_ms,rpc
	fmt.Printf("Go: github.com/scrapli/scrapligo (%s),%s,%d,%d,%d,%d,%.3f\n",
		transportType,
		framing,
		size,
		count,
		setupDur.Milliseconds(),
		duration.Milliseconds(),
		float64(count)/duration.Seconds())
}

func run(driver *netconf.Driver, size, count int) {
	filter := fmt.Sprintf(`<size>%d</size>`, size)
	for i := range count {
		logger := slog.With("req_id", i)

		logger.Info("sending request", "size", size)
		_, err := driver.Get(filter)
		if err != nil {
			logger.Error("request failed", "error", err)
		}
		logger.Info("request completed")
	}
}
