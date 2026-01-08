package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/driver/options"
)

var (
	host  = flag.String("host", "localhost", "NETCONF server host")
	port  = flag.Int("port", 8830, "NETCONF server port")
	user  = flag.String("user", "admin", "SSH username")
	pass  = flag.String("pass", "admin", "SSH password")
	size  = flag.Int("size", 1024, "Response size in bytes")
	count = flag.Int("count", 10, "Number of requests")
)

func main() {
	flag.Parse()

	target := fmt.Sprintf("%s:%d", *host, *port)
	slog.Info("connected", "server", target)

	driver, err := netconf.NewDriver(
		*host,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername(*user),
		options.WithAuthPassword(*pass),
		options.WithSSHConfigFile(""),
		options.WithTransportType("standard"),
		options.WithPort(*port),
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

	slog.Info("running requests", "count", *count, "size", *size)

	start := time.Now()

	filter := fmt.Sprintf(`<size>%d</size>`, *size)
	for i := 0; i < *count; i++ {
		logger := slog.With("req_id", i)

		logger.Info("sending request")
		_, err := driver.Get(filter)
		if err != nil {
			logger.Error("request failed", "error", err)
		}
		logger.Info("request completed")
	}

	duration := time.Since(start)

	fmt.Printf("%.3f\n", duration.Seconds())

	throughput := float64(*count) / duration.Seconds()
	dataTransferred := float64(*count**size) / (1024 * 1024)
	slog.Info("completed", "duration", duration, "throughput", throughput, "dataTransferred", dataTransferred)
}
