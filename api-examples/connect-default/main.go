package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_url]
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_url]
}