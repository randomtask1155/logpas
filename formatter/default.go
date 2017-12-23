package formatter

import (
	"io"
	"bufio"
	"fmt"
)

// default formatter does not reformat string. It will just pass it through
type DefaultFormat struct {
	Reader io.Reader
	Writer io.Writer
}

// Given a log line marsh
func (f DefaultFormat) Reformat(l []byte) ([]byte, error) {return l, nil}

// given a log line write it to the iowriter
func( f DefaultFormat) Write(l []byte) error {
	_, err  := fmt.Fprint(f.Writer, fmt.Sprintf("%s",l))
	if err != nil {
		return err
	}
	return nil
}

// Read next line from reader.  Users might want to override in case of multi line strings
func (f DefaultFormat) Read() error {
	r := bufio.NewReader(f.Reader)

	for {
		b, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return  err
		} else if err == io.EOF {
			break
		}
		l, err := f.Reformat(b)
		if err != nil {
			return err
		}
		f.Write(l)
	}
	return nil
}