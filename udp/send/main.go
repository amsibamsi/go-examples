// Package main sends all arguments as single packets via UDP to a destination.
package main

import (
	"flag"
	"log"
	"net"
)

var (
	address = flag.String("address", "localhost:8888", "Destination address as accepted by net.ResolveUDPAddr() in Golang")
)

func main() {
	flag.Parse()

	raddr, err := net.ResolveUDPAddr("udp", *address)
	if err != nil {
		log.Fatalf("Failed to resolve remote address: %v", err)
	}
	log.Printf("Remote address: %v", raddr)
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	log.Print("Connected ...")

	for i := 0; i < len(flag.Args()); i++ {
		data := []byte(flag.Arg(i))
		size := len(data)
		for offset := 0; offset < size; {
			n, err := conn.Write(data[offset:])
			if n > 0 {
				log.Printf("Sent data: %s", data[offset:n])
			}
			if err != nil {
				log.Printf("Error sending data: %v", err)
				// Prevent looping forever if no progress is made
				if n == 0 {
					log.Printf("Skipping data: %s", data[offset:])
					offset = len(data) - 1
				}
			}
			offset += n
		}
	}
}
