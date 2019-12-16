package errors

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/kamiaka/go-errors/internal/stack"
	"github.com/kamiaka/go-errors/lv"
)

// Error is typed error.
type Error interface {
	Error() string
	Unwrap() error
	Type() Type
	With(...*lv.LV) Error
}

type err struct {
	typ    Type
	msg    string
	frames *stack.Frames
	err    error
	lvs    []*lv.LV
}

func (e *err) Error() string {
	return e.msg
}

func (e *err) Unwrap() error {
	return e.err
}

func (e *err) Type() Type {
	return e.typ
}

func (e *err) Is(target error) bool {
	return e == target || e.typ == target
}

func (e *err) Format(s fmt.State, verb rune) {
	xerrors.FormatError(e, s, verb)
}

func (e *err) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.Error())
	if p.Detail() {
		e.typ.FormatError(p)
		p.Print("callers: ")
		e.frames.Format(p)
		for _, lv := range e.lvs {
			p.Print(lv, "\n")
		}
	}

	return e.err
}

func (e *err) With(lvs ...*lv.LV) Error {
	e.lvs = append(e.lvs, lvs...)
	return e
}
