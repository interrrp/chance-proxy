package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/lmittmann/tint"
)

func main() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{}),
	))

	server, err := net.Listen("tcp", address)
	if err != nil {
		slog.Error("failed to start server", "err", err)
		os.Exit(1)
	}
	defer server.Close()

	slog.Info("starting", "address", address, "target", target, "failureChance", failureChance)

	for {
		client, err := server.Accept()
		if err != nil {
			slog.Error("failed to accept connection", "err", err)
			continue
		}
		go handleClient(client)
	}
}
