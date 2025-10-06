package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	startA := time.Now()
	if _, err := http.Get(a); err != nil {
		panic(err)
	}
	timeA := time.Since(startA)

	startB := time.Now()
	if _, err := http.Get(b); err != nil {
		panic(err)
	}
	timeB := time.Since(startB)

	if timeA < timeB {
		return a
	}

	return b
}
