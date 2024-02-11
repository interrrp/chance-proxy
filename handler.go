package main

import (
	"io"
	"log"
	"math/rand"
	"net"
)

func handleClient(client net.Conn) {
	defer client.Close()

	if rand.Intn(100) < chance {
		log.Printf("disconnecting %s", client.RemoteAddr())
		return
	}

	log.Printf("proxying %s", client.RemoteAddr())

	svr, err := net.Dial("tcp", target)
	if err != nil {
		log.Printf("error dialing server: %v", err)
		return
	}
	defer svr.Close()

	go io.Copy(svr, client)
	io.Copy(client, svr)
}
