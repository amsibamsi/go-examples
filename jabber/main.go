package main

import (
	"github.com/mattn/go-xmpp"
	"github.com/vaughan0/go-ini"
	"log"
)

type Config struct {
	file ini.File
}

func NewConfig(filename string) Config {
	file, err := ini.LoadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return Config{file}
}

func (config Config) Get(section string, key string) string {
	val, ok := config.file.Get(section, key)
	if !ok {
		log.Fatalf("Not found in config: Key '%s' in section '%s'", key, section)
	}
	return val
}

func main() {
	config := NewConfig("config.ini")
	host := config.Get("server", "host")
	user := config.Get("login", "user")
	passwd := config.Get("login", "passwd")
	options := xmpp.Options{
		Host:                         host,
		User:                         user,
		Password:                     passwd,
		Resource:                     "bot",
		InsecureAllowUnencryptedAuth: false,
		NoTLS:         true,
		StartTLS:      true,
		Debug:         true,
		StatusMessage: "Alive and Kickin'",
	}
	client, err := options.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	for quit := false; !quit; {
		stanza, err := client.Recv()
		if err != nil {
			log.Fatal(err)
		}
		switch stanza.(type) {
		case xmpp.Presence:
			pres := stanza.(xmpp.Presence)
			log.Printf("presence: %+v", pres)
		case xmpp.Chat:
			chat := stanza.(xmpp.Chat)
			log.Printf("chat: %+v", chat)
			if chat.Text == "q" {
				quit = true
			}
		}
	}
}
