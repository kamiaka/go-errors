package errors

import (
	"fmt"
	"time"
)

// Error is error interface.
// that has WithFileds method.
type Error interface {
	Error() string
	WithFields() ErrorWithFields
}

type errorString struct {
	msg string
}

// New returns new error.
func New(msg string) Error {
	return &errorString{
		msg: msg,
	}
}

func (e *errorString) Error() string {
	return e.msg
}

func (e *errorString) WithFields() ErrorWithFields {
	return WithFields(e)
}

// WrappedError is error and wrapped error.
// that can handling original error and
type WrappedError interface {
	Error() string
	Origin() error
	WithFields() ErrorWithFields
}

type wrapError struct {
	error
	msg string
}

// Wrap error and returns error with fields.
func Wrap(err error, msg string) WrappedError {
	return &wrapError{
		error: err,
		msg:   msg,
	}
}

// Wrapf returns wrapped error.
func Wrapf(msg string, err error) WrappedError {
	return &wrapError{
		error: err,
		msg:   fmt.Sprintf(msg, err),
	}
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Origin() error {
	return e.error
}

func (e *wrapError) WithFields() ErrorWithFields {
	return WithFields(e)
}

// ErrorWithFields annotate
type ErrorWithFields interface {
	// Error is method for implementation of error.
	Error() string

	// Origin returns original error.
	Origin() error

	// Format is implementation for fmt.Formatter.
	Format(s fmt.State, verb rune)

	// Set Fields.
	Bool(label string, value bool) ErrorWithFields
	String(label string, value string) ErrorWithFields
	Stringer(label string, value fmt.Stringer) ErrorWithFields
	Bytes(label string, value []byte) ErrorWithFields
	Byte(label string, value byte) ErrorWithFields
	Int(label string, value int) ErrorWithFields
	Int64(label string, value int64) ErrorWithFields
	Int32(label string, value int32) ErrorWithFields
	Int16(label string, value int16) ErrorWithFields
	Int8(label string, value int8) ErrorWithFields
	Uint(label string, value uint) ErrorWithFields
	Uint64(label string, value uint64) ErrorWithFields
	Uint32(label string, value uint32) ErrorWithFields
	Uint16(label string, value uint16) ErrorWithFields
	Uint8(label string, value uint8) ErrorWithFields
	Float64(label string, value float64) ErrorWithFields
	Float32(label string, value float32) ErrorWithFields
	Time(label string, value time.Time, layout string) ErrorWithFields
	UTCTime(label string, value time.Time) ErrorWithFields
	Stack(label string) ErrorWithFields
}
