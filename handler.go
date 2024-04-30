package main

import (
	"io"
	"log/slog"
	"math/rand"
	"net"
)

func handleClient(client net.Conn) {
	defer client.Close()

	if rand.Intn(100) < chance {
		slog.Info("disconnecting", "ip", client.RemoteAddr())
		return
	}

	slog.Info("proxying", "ip", client.RemoteAddr())

	target, err := net.Dial("tcp", target)
	if err != nil {
		slog.Error("failed to dial server", "err", err)
		return
	}
	defer target.Close()

	go io.Copy(target, client)
	io.Copy(client, target)
}
