package errors

import (
	"fmt"
	"io"
)

type field struct {
	key   string
	value interface{}
}

func (f *field) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, f.key)
			io.WriteString(s, ":")
			fmt.Fprintf(s, "%+v", f.value)
		}
	}
}

type fields []field

func (b *fields) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			for _, f := range *b {
				fmt.Fprintf(s, "\n%+v", f)
			}
		}
	}
}
