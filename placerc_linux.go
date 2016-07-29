package main

import (
	"log"
	"os"
)

const logPath = "/var/log/cdrserver.log"

func init() {
	if err := os.Setenv("CDRSERVERRC", "/usr/local/etc/cdrserverrc"); err != nil {
		log.Println(err)
	}
}
