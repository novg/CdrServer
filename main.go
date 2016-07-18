package main

import (
	"log"
	"os"

	"novg/cdrserver/server"
)

func main() {
	f, err := os.OpenFile("listen.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	// port := flag.Int("port", 2020, "`PORT` for listening of CDR clients")
	// fmt.Printf("port: %d\n", *port)
	server.Run()
}
