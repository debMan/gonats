package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	natsURL        = "nats://127.0.0.1:4222"
	natsUser       = "admin"
	natsPass       = "password"
	sizeKB         = 1024
	streamName     = "rides"
	subjectName    = ">" // Or use ">" if you want to catch all
	ackDurationSec = 30
	maxAckPending  = 1024
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  ./gonats publish <message_count>")
		fmt.Println("  ./gonats consume <worker_count>; 0 to infinite in produce")
		return
	}

	mode := os.Args[1]
	countStr := os.Args[2]

	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Println("you must specify a valid number for worker/message count; 0 for infinite messages produce")
		os.Exit(1)
	}

	switch mode {
	case "publish":
		Publish(sizeKB, count)
	case "consume":
		Consume(count)
	default:
		log.Println("you must specify publish or consume mode")
	}
}
