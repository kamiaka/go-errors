package errors

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// New returns new error.
func New(msg string) error {
	return errors.New(msg)
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
