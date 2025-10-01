package main

import (
	"log"
	"os"

	"github.com/bingbeann/go-tdd/countdown"
)

func main() {
	if err := countdown.Countdown(os.Stdout); err != nil {
		log.Fatalf("failed to countdown: %v", err)
	}
}
