package main

import "flag"

var (
	address string
	target  string
	chance  int
)

func init() {
	flag.StringVar(&address, "address", "localhost:8081", "Address to listen on")
	flag.StringVar(&target, "target", "localhost:8080", "Target address")
	flag.IntVar(&chance, "chance", 20, "Chance of failure")
	flag.Parse()
}
