package main

import (
	"log"
	"os"
	"time"

	"github.com/bingbeann/go-tdd/countdown"
)

func main() {
	if err := countdown.Countdown(os.Stdout, 3, &DefaultSleeper{}); err != nil {
		log.Fatalf("failed to countdown: %v", err)
	}
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep(d time.Duration) {
	time.Sleep(d)
}
