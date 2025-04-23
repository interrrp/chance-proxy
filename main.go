package main

import (
	"io"
	"log/slog"
	"math/rand"
	"net"
	"os"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("failed to load .env", "err", err)
	}

	cfg, err := readConfig()
	must(err, "failed to read config")

	srv := server{cfg}
	slog.Info("starting server", "cfg", cfg)
	must(srv.start(), "failed to start server")
}

func must(err error, msg string) {
	if err != nil {
		slog.Error(msg, "err", err)
		os.Exit(1)
	}
}

type config struct {
	Address       string `env:"ADDRESS"`
	Target        string `env:"TARGET"`
	FailureChance int    `env:"FAILURE_CHANCE"`
}

func readConfig() (config, error) {
	opts := env.Options{
		Prefix: "CPXY_",
	}
	return env.ParseAsWithOptions[config](opts)
}

type server struct {
	cfg config
}

func (s *server) start() error {
	ln, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		return err
	}

	for {
		client, err := ln.Accept()
		if err != nil {
			slog.Error("failed to accept connection", "err", err)
			continue
		}

		go s.handleClient(client)
	}
}

func (s *server) handleClient(client net.Conn) {
	defer client.Close()

	if rand.Intn(100) < s.cfg.FailureChance {
		slog.Info("disconnecting", "ip", client.RemoteAddr())
		return
	}

	slog.Info("proxying", "ip", client.RemoteAddr())

	target, err := net.Dial("tcp", s.cfg.Target)
	if err != nil {
		slog.Error("failed to dial target", "err", err, "target", s.cfg.Target)
		return
	}
	defer target.Close()

	proxyConnections(target, client)
}

func proxyConnections(target, client net.Conn) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if _, err := io.Copy(target, client); err != nil {
			slog.Error("client to target copy failed", "err", err)
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := io.Copy(client, target); err != nil {
			slog.Error("target to client copy failed", "err", err)
		}
	}()

	wg.Wait()
}
