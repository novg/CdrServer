package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"novg/cdrserver/dbclient"
	"novg/cdrserver/server"
	"os"

	_ "github.com/lib/pq"
	"github.com/schachmat/ingo"
)

var (
	port     *int
	host     *string
	name     *string
	user     *string
	password *string
)

func main() {
	f, err := os.OpenFile("cdrserver.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	initSettings()

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		*host, *user, *password, *name)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	dbclient.InitDatabase(&db)
	server.Run(*port)
}

func initSettings() {
	port = flag.Int("cport", 2112, "`PORT` for listening of CDR clients (PBX)")
	host = flag.String("host", "localhost", "`DB_HOST` database")
	name = flag.String("name", "cdrbase", "`DB_NAME` is name of database")
	user = flag.String("user", "aastra", "`DB_USER` is user of database")
	password = flag.String("password", "aastra", "`DB_PASSWORD` is password of database")
	// if err := os.Setenv("CDRSERVERRC", "cdrserverrc"); err != nil {
	// 	log.Println(err)
	// }
	if err := ingo.Parse("cdrserver"); err != nil {
		log.Println(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
