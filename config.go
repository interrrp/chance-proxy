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
	godotenv.Load()
	address = getEnv("CPXY_ADDRESS", ":8081")
	target = getEnv("CPXY_TARGET", ":8080")
	failureChance = getIntEnv("CPXY_FAILURE_CHANCE", 20)

	flag.StringVar(&address, "address", address, "address to listen on")
	flag.StringVar(&target, "target", target, "address to proxy to")
	flag.IntVar(&failureChance, "failure-chance", failureChance, "percent chance of failure")
	flag.Parse()
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
