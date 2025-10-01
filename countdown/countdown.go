package countdown

import (
	"fmt"
	"io"
)

func Countdown(w io.Writer) error {
	if _, err := fmt.Fprint(w, "3"); err != nil {
		return err
	}

	return nil
}
