package racer

import (
	"errors"
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("reached timeout of 10 seconds")
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
