package main

import (
	"log"
	"os"
)

func init() {
	if err := os.Setenv("CDRSERVERRC", "cdrserverrc"); err != nil {
		log.Println(err)
	}
}
