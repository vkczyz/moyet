package main

import (
	"bufio"
	"crypto/tls"
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

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		//data := scanner.Text()

		status := "20"
		meta := "text/gemini; charset=utf-8"
		body := "This is a test response"
		conn.Write([]byte(status + " " + meta + "\r\n" + body))
	}
	if scanner.Err() != nil {
		log.Print("Could not scan input")
	}
}