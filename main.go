package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net"
)

const (
	IP = "0.0.0.0"
	PORT = "1965"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	address := IP + ":" + PORT
	listener, err := tls.Listen("tcp", address, config)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		log.Print("Connection established")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	status := "20"
	meta := "text/gemini; charset=utf-8"

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		status = "50"
		meta = "Unreadable input"
	}

	body, err := ioutil.ReadFile("example.gmi")
	if err != nil {
		status = "40"
		meta = "Could not find the requested file"
	}

	response := []byte(status + " " + meta + "\r\n")
	response = append(response, body...)
	conn.Write(response)
}