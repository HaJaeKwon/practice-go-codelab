package main

import (
	"codelab/pkg/chat"
	"log"
)

func main() {
	server, err := chat.NewServer(":8080")
	if err != nil {
		log.Fatalln(err)
	}

	err = server.Listen()
	if err != nil {
		log.Fatalln(err)
	}
}
