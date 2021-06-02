package main

import (
	"log"

	"github.com/AlfredoPastor/ari-queue/internal"
)

func main() {
	server, err := internal.InitializeServer()
	if err != nil {
		log.Fatal("error compilation time: ", err.Error())
	}
	if err := server.Run(); err != nil {
		log.Fatal("error execution time: ", err.Error())
	}
}
