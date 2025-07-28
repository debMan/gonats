package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

func Consume(workers int) {
	nc, err := nats.Connect(natsURL, nats.UserInfo(natsUser, natsPass))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for i := range workers {
		wg.Add(1)
		queueGroup := fmt.Sprintf("durable-worker-%d", i)
		go func(id int) {
			defer wg.Done()

			_, err := js.QueueSubscribe(subjectName, queueGroup,
				func(msg *nats.Msg) {
					fmt.Printf("[Worker %d] Got message (%d bytes): %.10s...\n", id, len(msg.Data), string(msg.Data))
					msg.Ack()
				}, nats.Durable(fmt.Sprintf("durable-worker-%d", id)),
				nats.ManualAck(),
				nats.AckWait(ackDurationSec*time.Second),
				nats.MaxAckPending(maxAckPending),
			)
			if err != nil {
				log.Printf("[Worker %d] Subscription error: %v", id, err)
				return
			}
			fmt.Printf("[Worker %d] Listening...\n", id)

			// Keep the goroutine alive
			select {}
		}(i)
	}

	wg.Wait()
}
