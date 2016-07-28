package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"novg/cdrserver/dbclient"
)

// Run launch listening server on port
func Run(port int) {
	// Listen on TCP port on all interfaces.
	localPort := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp4", localPort)
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
				return
			}
		}

		// send message to database
		fmt.Fprint(dbclient.CallInfo, message)
	}
}
