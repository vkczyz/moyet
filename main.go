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
	ROOT = "."
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
	body := []byte("")

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		status = "59"
		meta = "Could not read input"
	}
	
	requested, err := NewRequest(buf)
	if err != nil {
		status = "59"
		meta = "Could not parse URL"
	} else {
		path := requested.url.RequestURI()
		log.Print(path)

		filepath := ROOT + path
		if filepath[len(filepath)-1:] == "/" {
			filepath += "index.gmi"
		}

		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			status = "51"
			meta = "Could not find the requested file"
		} else {
			body = data
		}
	}

	response := []byte(status + " " + meta + "\r\n")
	response = append(response, body...)
	conn.Write(response)
}