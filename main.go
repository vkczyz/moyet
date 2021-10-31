package main

import (
	"bufio"
	"log"
	"net"
)

const (
	IP = "0.0.0.0"
	PORT = "1965"
)

func main() {
	address := IP + ":" + PORT
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

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