package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Config struct {
	Login struct {
		User   string
		Passwd string
	}
	Server struct {
		Host string
		Port int
	}
}

func main() {
	const jdata = `
{
  "login": {
    "user": "bob",
    "passwd": "1234"
  },
	"server": {
		"host": "localhost",
		"port": 2345
	}
}

`
	dec := json.NewDecoder(strings.NewReader(jdata))
	var conf Config
	err := dec.Decode(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", conf)
}
