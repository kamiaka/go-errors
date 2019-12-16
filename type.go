package errors

import (
	"fmt"
	"strings"

	"golang.org/x/xerrors"

	"github.com/kamiaka/go-errors/internal/stack"
)

// Type is error type.
type Type interface {
	// Error method for implements build-in error interface.
	Error() string

	// New returns an Error that formats as the given text.
	New(text string) Error

	// Errorf formats according to a format specifier
	// and returns the string as a value that satisfies error.
	//
	// Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, a ...interface{}) Error

	// Wrap error.
	Wrap(e error) Error

	Format(f fmt.State, verb rune)

	// FormatError implements xerrors.Formatter.
	// FormatError prints the receiver's first error and returns the next error in
	// the error chain, if any.
	FormatError(p xerrors.Printer) (next error)
}

type typ struct {
	name  string
	level string
	skip  int
	depth int
}

// NewType returns new error type.
func NewType(name string, opts ...TypeOption) Type {
	t := &typ{
		name:  name,
		level: "error",
		skip:  0,
		depth: 1,
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func (t *typ) Error() string {
	return "error type: " + t.name
}

func (t *typ) new(msg string, origin error) Error {
	return &err{
		typ:    t,
		msg:    msg,
		frames: stack.Callers(t.depth, t.skip+2),
		err:    origin,
	}
}

func (t *typ) New(text string) Error {
	return t.new(text, nil)
}

func (t *typ) Errorf(format string, a ...interface{}) Error {
	var err error

	if len(a) > 0 {
		if e, ok := a[len(a)-1].(error); ok && strings.HasSuffix(format, ": %w") {
			format = strings.TrimSuffix(format, ": %w") + ": %v"
			err = e
		}
	}

	return t.new(fmt.Sprintf(format, a...), err)
}

func (t *typ) Wrap(err error) Error {
	return t.new(err.Error(), err)
}

func (t *typ) Level() string {
	return t.level
}

func (t *typ) Format(s fmt.State, verb rune) {
	xerrors.FormatError(t, s, verb)
}

func (t *typ) FormatError(p xerrors.Printer) (next error) {
	p.Print(t.Error())
	if p.Detail() {
		p.Printf(", level: %s\n", t.level)
	}
	return nil
}
