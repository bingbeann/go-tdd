package countdown

import (
	"fmt"
	"io"
)

func Countdown(w io.Writer, n int) error {
	for i := n; i >= 1; i-- {
		if _, err := fmt.Fprintf(w, "%d\n", i); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "Go!"); err != nil {
		return err
	}

	return nil
}
