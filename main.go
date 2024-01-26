package main

import (
	"log"
	"net"
	"strings"
)

type TCPProtocol struct {
	protocol string
	addr     string
}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func handleConnection(conn net.Conn) {
	var buff []byte = make([]byte, 4098)
	_, err := conn.Read(buff)

	request := string(buff)

  strarr := strings.Split(request, "\r\n")

  for index, data := range strarr {
    log.Println(index, " ", data)
  }

	var writeBuff []byte = []byte("HTTP/1.1 200 OK\r\n")
  _, err = conn.Write(writeBuff)
  handleError(err)
  
  conn.Close()
}

func main() {
	log.Println("Starting server...")

	var httpProtocol = TCPProtocol{protocol: "tcp", addr: "127.0.0.1:8080"}
	listener, err := net.Listen(httpProtocol.protocol, httpProtocol.addr)

	handleError(err)

	for {
		conn, err := listener.Accept()
		handleError(err)
		go handleConnection(conn)
		defer conn.Close()
	}

}
