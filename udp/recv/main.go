// Package main receives data via UDP on a local address and prints it out until
// it receives a SIGTERM or receives the byte sequence "end".
package main

import (
	"flag"
	"log"
	"net"
)

var (
	address = flag.String("address", "localhost:8888", "Listen address as accepted by net.ResolveUDPAddr() in Golang")
)

func main() {
	flag.Parse()

	laddr, err := net.ResolveUDPAddr("udp", *address)
	if err != nil {
		log.Fatalf("Failed to resolve listen addr: %v", err)
	}
	log.Printf("Listen address: %v", laddr)
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()
	log.Print("Listening ...")

	// Max UDP packet size is 2^16 bytes
	buf := make([]byte, 2<<15)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		log.Printf("Data from %v: %s", addr, buf[:n])
		if err != nil {
			log.Printf("Error reading: %v", err)
		}
	}
}
