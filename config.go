package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	address string
	target  string
	chance  int
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}
	address = getEnv("ADDRESS", ":8081")
	target = getEnv("TARGET", ":8080")
	chance = getIntEnv("CHANCE", 20)
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
