package errors

import (
	"runtime"
	"strconv"
	"strings"
)

func callers() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(4, pcs[:])

	return pcs[0:n]
}

func appendCallers(buf []byte) []byte {
	pcs := callers()
	for i, pc := range pcs {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			file, line := fn.FileLine(pc)
			name := fn.Name()
			if i != 0 {
				buf = append(buf, ',')
			}
			buf = append(buf, name...)
			buf = append(buf, '(', ')', ':')
			buf = append(buf, trimGOPATH(name, file)...)
			buf = append(buf, ':')
			buf = strconv.AppendInt(buf, int64(line), 10)
		}
	}

	return buf
}

// NOTE: copy from https://github.com/pkg/errors/blob/c605e284fe17294bda444b34710735b29d1a9d90/stack.go#L148
func trimGOPATH(name, file string) string {
	// Here we want to get the source file path relative to the compile time
	// GOPATH. As of Go 1.6.x there is no direct way to know the compiled
	// GOPATH at runtime, but we can infer the number of path segments in the
	// GOPATH. We note that fn.Name() returns the function name qualified by
	// the import path, which does not include the GOPATH. Thus we can trim
	// segments from the beginning of the file path until the number of path
	// separators remaining is one more than the number of path separators in
	// the function name. For example, given:
	//
	//    GOPATH     /home/user
	//    file       /home/user/src/pkg/sub/file.go
	//    fn.Name()  pkg/sub.Type.Method
	//
	// We want to produce:
	//
	//    pkg/sub/file.go
	//
	// From this we can easily see that fn.Name() has one less path separator
	// than our desired output. We count separators from the end of the file
	// path until it finds two more than in the function name and then move
	// one character forward to preserve the initial path segment without a
	// leading separator.
	const sep = "/"
	goal := strings.Count(name, sep) + 2
	i := len(file)
	for n := 0; n < goal; n++ {
		i = strings.LastIndex(file[:i], sep)
		if i == -1 {
			// not enough separators found, set i so that the slice expression
			// below leaves file unmodified
			i = -len(sep)
			break
		}
	}
	// get back to 0 or trim the leading separator
	file = file[i+len(sep):]
	return file
}
