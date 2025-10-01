package greet

import (
	"io"
)

func Greet(w io.Writer, name string) {
	w.Write([]byte("Hello, " + name))
}
