package ofmt

import "io"

type buffer []byte

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	n, err = w.Write([]byte(format))
	return
}