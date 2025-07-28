package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

func Publish(sizeKB, count int) {

	// Connect to NATS
	nc, err := nats.Connect(natsURL, nats.UserInfo(natsUser, natsPass))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// Generate a large payload
	payload := randomString(sizeKB)

	var wg sync.WaitGroup
	for _, subject := range subjects {
		// capture range variable correctly, keep it if you want backward compatibility with Go <1.22
		// subject := subject
		wg.Add(1)
		go func() {
			defer wg.Done()
			// for infinite loop
			if count == 0 {
				fmt.Printf("Publishing messages to subject: %s\n", subject)
				for {
					if _, err := js.Publish(subject, []byte(payload)); err != nil {
						log.Printf("Error publishing to %s: %v", subject, err)
					}
				}
			} else {
				fmt.Printf("Publishing %d messages to subject: %s\n", count, subject)
				for range count {
					if _, err := js.Publish(subject, []byte(payload)); err != nil {
						log.Printf("Error publishing to %s: %v", subject, err)
					}
				}
			}
		}()
	}

	wg.Wait()
	fmt.Println("Done.")
}
