package stack

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"

	"golang.org/x/xerrors"
)

// Frames ...
type Frames struct {
	*runtime.Frames
}

// Callers ...
func Callers(depth, skip int) *Frames {
	pc := make([]uintptr, depth)

	runtime.Callers(skip+2, pc)

	fmt.Printf("callers: %#v, depth: %d, skip: %d\n", pc, depth, skip+2)

	return &Frames{
		runtime.CallersFrames(pc),
	}
}

// Format prints the stack as error detail.
func (f *Frames) Format(p xerrors.Printer) {
	if p.Detail() {
		p.Print(f.String())
	}
}

// String
func (f *Frames) String() string {
	var b bytes.Buffer
	inMore := false
	for {
		frame, more := f.Next()
		if frame.Function == "runtime.main" {
			break
		}
		if inMore {
			b.WriteString("\n  ")
		}
		b.WriteString(frame.Function + "\n")
		b.WriteString("    " + frame.File + ":" + strconv.Itoa(frame.Line))
		if more {
			inMore = true
			continue
		}
		break
	}

	return b.String()
}
