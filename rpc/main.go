package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Service struct{}

func (s *Service) Function(req []byte, resp *[]byte) error {
	var n int64
	if err := binary.Read(bytes.NewReader(req), binary.BigEndian, &n); err != nil {
		return err
	}
	n++
	b := new(bytes.Buffer)
	if err := binary.Write(b, binary.BigEndian, n); err != nil {
		return err
	}
	*resp = b.Bytes()
	return nil
}

func run() error {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	defer l.Close()
	if err != nil {
		return err
	}
	s := rpc.NewServer()
	s.Register(new(Service))
	go func() {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
		}
		s.ServeConn(conn)
	}()
	client, err := rpc.Dial(l.Addr().Network(), l.Addr().String())
	if err != nil {
		return err
	}
	var n int64
	n = 42
	req := new(bytes.Buffer)
	if err := binary.Write(req, binary.BigEndian, n); err != nil {
		return err
	}
	var resp []byte
	if err := client.Call("Service.Function", req.Bytes(), &resp); err != nil {
		return err
	}
	if err := binary.Read(bytes.NewReader(resp), binary.BigEndian, &n); err != nil {
		return err
	}
	fmt.Println(n)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
