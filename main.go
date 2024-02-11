package main

import (
	"log"
	"net"
)

func main() {
	s, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
	defer s.Close()

	log.Printf("listening on %s", address)
	log.Printf("proxying to %s", target)
	log.Printf("chance of failure: %d%%", chance)

	for {
		client, err := s.Accept()
		if err != nil {
			log.Printf("error accepting connection: %v", err)
			continue
		}
		go handleClient(client)
	}
}
