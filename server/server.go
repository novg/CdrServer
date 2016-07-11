package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Run launch listening server
func Run() {
	// Listen on TCP port 2112 on all interfaces.
	ln, err := net.Listen("tcp4", ":2112")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		// Wait for a connection.
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiplay connections may be servered concurently.
		// go handleConnection(conn)
		// will listen for message to process ending in newline (\n)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	client := conn.RemoteAddr()
	defer conn.Close()

	log.Printf("%s is connected\r\n", client)
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			switch err {
			case io.EOF:
				log.Printf("%s is disconnected: %v\r\n", client, err)
				return
			default:
				log.Printf("bad message: %v\r\n", err)
			}
		}

		// TODO: send to database
		f, err := os.OpenFile("out.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		//
		fmt.Fprint(f, message)
	}
}
