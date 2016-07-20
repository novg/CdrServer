package main

import (
	"flag"
	"log"
	"novg/cdrserver/server"
	"os"

	"github.com/novg/ingo"
)

var (
	port *int
)

func main() {
	f, err := os.OpenFile("listen.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	parseSettings()

	server.Run(*port)
}

func parseSettings() {
	port = flag.Int("port", 2112, "`PORT` for listening of CDR clients (PBX)")
	host := flag.String("host", "localhost", "`HOST` database")
	base := flag.String("base", "calls_database", "`BASE` is name of database")
	user := flag.String("user", "aastra", "`USER` is user of database")
	password := flag.String("password", "aastra", "`PASSWORD` is password of database")
	if err := os.Setenv("CDRSERVERRC", "cdrserverrc"); err != nil {
		log.Println(err)
	}
	if err := ingo.Parse("cdrserver"); err != nil {
		log.Println(err)
	}

	// TODO: delete
	log.Println(*port, *host, *base, *user, *password)
}
