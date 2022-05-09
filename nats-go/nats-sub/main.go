package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"os"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe
	args := os.Args
	if len(args) != 2 { return }

	subj, i := args[1], 0

	nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		log.Printf("[#%d] Received on [%s]: '%s'", i, msg.Subject, string(msg.Data))
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", subj)
	runtime.Goexit()
}

