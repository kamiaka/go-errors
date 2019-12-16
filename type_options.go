package errors

// TypeOption is option for new error type.
type TypeOption func(*typ)

// Level set error level to error type.
func Level(level string) TypeOption {
	return func(t *typ) {
		t.level = level
	}
}

// Skip `n` frame when get caller.
func Skip(n int) TypeOption {
	return func(t *typ) {
		t.skip = n + 1
	}
}

// Depth gets `n` layer stack trace.
func Depth(n int) TypeOption {
	return func(t *typ) {
		t.depth = n
	}
}
