package countdown

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep(time.Duration)
}

func Countdown(w io.Writer, n int, sleeper Sleeper) error {
	for i := n; i >= 1; i-- {
		if _, err := fmt.Fprintln(w, i); err != nil {
			return err
		}

		sleeper.Sleep(1 * time.Second)
	}

	if _, err := fmt.Fprint(w, "Go!"); err != nil {
		return err
	}

	return nil
}
