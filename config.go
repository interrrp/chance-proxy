package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	address       string
	target        string
	failureChance int
)

func init() {
	flag.StringVar(&address, "address", ":8081", "address to listen on")
	flag.StringVar(&target, "target", ":8080", "address to proxy to")
	flag.IntVar(&failureChance, "failure-chance", 20, "percent chance of failure")
	flag.Parse()

	godotenv.Load()
	address = getEnv("CPXY_ADDRESS", address)
	target = getEnv("CPXY_TARGET", target)
	failureChance = getIntEnv("CPXY_FAILURE_CHANCE", failureChance)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getIntEnv(key string, fallback int) int {
	i, err := strconv.Atoi(getEnv(key, strconv.Itoa(fallback)))
	if err != nil {
		return fallback
	}
	return i
}
