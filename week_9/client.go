package main

import (
	"github.com/Terry-Mao/goim/api/protocol"
	"github.com/Terry-Mao/goim/pkg/bytes"
	"log"
	"net"
)

func RunClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Printf("connet server failed: %v", err)
		return
	}
	defer conn.Close()
	log.Printf("start to send msg")

	body := "this is client"
	writer := bytes.NewWriterSize(len(body) + 64)
	proto := &protocol.Proto{
		Ver:  1,
		Op:   2,
		Seq:  3,
		Body: []byte(body),
	}
	proto.WriteTo(writer)

	_, err = conn.Write(writer.Buffer())
	if err != nil {
		log.Printf("write msg failed: %v", err)
		return
	}
}
