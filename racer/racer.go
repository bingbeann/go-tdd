package racer

import (
	"net/http"
)

func Racer(a, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		if _, err := http.Get(url); err != nil {
			panic(err)
		}

		close(ch)
	}()
	return ch
}
