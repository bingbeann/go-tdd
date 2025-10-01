package greet

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, name string) error {
	if _, err := fmt.Fprintf(w, "Hello, %s", name); err != nil {
		return err
	}

	return nil
}
