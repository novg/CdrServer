package main

import (
	"log"
	"novg/cdrserver/server"
	"os"
)

func main() {
	f, err := os.OpenFile("listen.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	server.Run()
}
