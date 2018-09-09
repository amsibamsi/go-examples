package main

import (
	"log"
)

func main() {
	log.Printf("CPU Vendor: %s", cpuVendor)
	log.Printf("CPU AVX2 Support: %t", cpuAVX2)
}
