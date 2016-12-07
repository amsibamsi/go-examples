package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("main ")
	log.Print("message1")
	l := log.New(os.Stdout, "sub ", log.Lshortfile)
	l.Printf("message2")
}
