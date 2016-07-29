package main

import (
	"log"
	"os"
)

const logPath = "cdrserver.log"

func init() {
	if err := os.Setenv("CDRSERVERRC", "cdrserverrc"); err != nil {
		log.Println(err)
	}
}
