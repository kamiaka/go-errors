package errors

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// WithFields annotate err with fields.
func WithFields(err error) ErrorWithFields {
	e, ok := err.(*withBuffer)
	if !ok {
		e = &withBuffer{
			error: err,
			buf:   []byte{},
		}
	}
	return e
}

type withBuffer struct {
	error
	buf []byte
}

func (b *withBuffer) Error() string {
	return b.error.Error()
}

func (b *withBuffer) Origin() error {
	return b.error
}

func (b *withBuffer) Unwrap() error {
	return b.error
}

func (b *withBuffer) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, b.error.Error())
			io.WriteString(s, "\n")
			s.Write(b.buf)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, b.error.Error())
	case 'q':
		fmt.Fprintf(s, "%q", b.error.Error())
	}
}

func (b *withBuffer) String(label string, value string) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = append(b.buf, escape(value)...)
	return b
}

func (b *withBuffer) Stringer(label string, value fmt.Stringer) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = append(b.buf, escape(value.String())...)
	b.buf = append(b.buf, '"')
	return b
}
func (b *withBuffer) Bytes(label string, value []byte) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = appendHexBytes(b.buf, value)
	b.buf = append(b.buf, '"')
	return b
}
func (b *withBuffer) Byte(label string, value byte) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = appendHexByte(b.buf, value)
	b.buf = append(b.buf, '"')
	return b
}

func (b *withBuffer) Bool(label string, value bool) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendBool(b.buf, value)
	return b
}

func (b *withBuffer) Int64(label string, value int64) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendInt(b.buf, value, 10)
	return b
}

func (b *withBuffer) Int32(label string, value int32) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendInt(b.buf, int64(value), 10)
	return b
}

func (b *withBuffer) Int16(label string, value int16) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendInt(b.buf, int64(value), 10)
	return b
}

func (b *withBuffer) Int8(label string, value int8) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendInt(b.buf, int64(value), 10)
	return b
}

func (b *withBuffer) Int(label string, value int) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendInt(b.buf, int64(value), 10)
	return b
}

func (b *withBuffer) Uint64(label string, value uint64) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendUint(b.buf, value, 10)
	return b
}

func (b *withBuffer) Uint32(label string, value uint32) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendUint(b.buf, uint64(value), 10)
	return b
}

func (b *withBuffer) Uint16(label string, value uint16) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendUint(b.buf, uint64(value), 10)
	return b
}

func (b *withBuffer) Uint8(label string, value uint8) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendUint(b.buf, uint64(value), 10)
	return b
}

func (b *withBuffer) Uint(label string, value uint) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendUint(b.buf, uint64(value), 10)
	return b
}

func (b *withBuffer) Float64(label string, value float64) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendFloat(b.buf, value, 'g', -1, 64)
	return b
}

func (b *withBuffer) Float32(label string, value float32) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = strconv.AppendFloat(b.buf, float64(value), 'g', -1, 32)
	return b
}

func (b *withBuffer) Time(label string, value time.Time, layout string) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	if layout == "" {
		layout = time.RFC3339
	}
	b.buf = append(b.buf, escape(value.Format(layout))...)
	return b
}

func (b *withBuffer) UTCTime(label string, value time.Time) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = appendUTCTime(b.buf, value)
	return b
}

func (b *withBuffer) Stack(label string) ErrorWithFields {
	b.buf = append(b.buf, '\t')
	b.buf = append(b.buf, label...)
	b.buf = append(b.buf, ':')
	b.buf = appendCallers(b.buf)
	return b
}
