package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"flag"
)

func main() {
	var reply = flag.String("reply", "", "Sets a specific reply subject")

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe
	args := os.Args
	if len(args) != 3 { return }

	subj, msg := args[1], []byte(args[2])

	if reply != nil && *reply != "" {
		nc.PublishRequest(subj, *reply, msg)
	} else {
		nc.Publish(subj, msg)
	}

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
	
}

