package wcl

import (
	"os"
	"io"
	"bufio"
)

func FromFile (fname string) (int, error) {
	nlines := 0
	fh, err := os.Open(fname)
	defer fh.Close()
	if err != nil {
		return -1, err
	}
	buf := bufio.NewReader(fh)
	for {
		_, _, err := buf.ReadLine()
		if err == io.EOF {
			return nlines, nil
		}
		nlines++
	}
	return nlines, nil // never used
}
