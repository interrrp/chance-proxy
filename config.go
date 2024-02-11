package main

import "flag"

var (
	address = *flag.String("address", "localhost:8081", "Address to listen on")
	target  = *flag.String("target", "localhost:8080", "Target address")
	chance  = *flag.Int("chance", 20, "Chance of failure")
)

func init() {
	flag.Parse()
}
