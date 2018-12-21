package main

import (
	"log"
	"io"
	"net"
	"time"
)

func main() {
	listener, err :=  net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("03:04:05\n"))
		if err != nil {
			panic(err)
		}
		time.Sleep(1*time.Second)
	}
}