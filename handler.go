package main

import "net"

func handleClient(client net.Conn) {
	defer client.Close()
}
