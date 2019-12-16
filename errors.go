package errors

var defaultError = NewType("Error!", Skip(1))

// New error.
func New(text string) Error {
	return defaultError.New(text)
}

// Errorf formats according to a format specifier
// and returns the string as a value that satisfies error.
func Errorf(format string, a ...interface{}) Error {
	return defaultError.Errorf(format, a...)
}
