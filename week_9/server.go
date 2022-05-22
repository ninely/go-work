package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

type Proto struct {
	PackLen   int32  // package length
	HeaderLen int16  // header length
	Ver       int16  // protocol version
	Operation int32  // operation for request
	Seq       int32  // sequence number chosen by client
	Body      []byte // body
}

func (p *Proto) GetBody() string {
	return string(p.Body)
}

func Listen() {
	socket, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Printf("listen failed: %v", err)
		return
	}
	defer socket.Close()
	log.Println("start listening")

	conn, err := socket.Accept()
	if err != nil {
		log.Printf("accept failed: %v", err)
		return
	}

	defer conn.Close()
	msg := new(Proto)
	if err := decode(conn, msg); err != nil {
		log.Printf("decode failed: %v", err)
		return
	}
	log.Printf("msg: %v", msg.GetBody())
}

func decode(rd io.Reader, proto *Proto) (err error) {
	var (
		packLen   int32
		headerLen int16
	)

	if err = binary.Read(rd, binary.BigEndian, &packLen); err != nil {
		return
	}
	if err = binary.Read(rd, binary.BigEndian, &headerLen); err != nil {
		return
	}
	if err = binary.Read(rd, binary.BigEndian, &proto.Ver); err != nil {
		return
	}
	if err = binary.Read(rd, binary.BigEndian, &proto.Operation); err != nil {
		return
	}
	if err = binary.Read(rd, binary.BigEndian, &proto.Seq); err != nil {
		return
	}
	var (
		n, t    int
		bodyLen = int(packLen - int32(headerLen))
	)
	if bodyLen > 0 {
		proto.Body = make([]byte, bodyLen)
		for {
			if t, err = rd.Read(proto.Body[n:]); err != nil {
				return
			}
			if n += t; n == bodyLen {
				break
			}
		}
	} else {
		proto.Body = nil
	}
	return
}
