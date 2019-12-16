package lv

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kamiaka/go-errors/internal/stack"
)

// LV is label and string value.
type LV struct {
	Label string
	Value string
}

// String returns label and value string.
func (v *LV) String() string {
	return fmt.Sprintf("%s: %s", v.Label, v.Value)
}

// String returns LV.
func String(l, v string) *LV {
	return &LV{
		Label: l,
		Value: v,
	}
}

// Stringer returns LV.
func Stringer(l string, v interface{ String() string }) *LV {
	return &LV{
		Label: l,
		Value: v.String(),
	}
}

// Bool returns LV.
func Bool(l string, v bool) *LV {
	var s string
	if v {
		s = "true"
	} else {
		s = "false"
	}
	return &LV{
		Label: l,
		Value: s,
	}
}

var hexPrefix = []byte("0x")
var digits = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

// Bytes returns LV.
func Bytes(l string, v []byte) *LV {
	buf := hexPrefix
	for _, b := range v {
		buf = append(buf, digits[b/16], digits[b%16])
	}
	return &LV{
		Label: l,
		Value: string(buf),
	}
}

// Byte returns LV.
func Byte(l string, b byte) *LV {
	return &LV{
		Label: l,
		Value: string(append(hexPrefix, digits[b/16], digits[b%16])),
	}
}

// Int returns LV.
func Int(l string, v int) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int8 returns LV.
func Int8(l string, v int8) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int16 returns LV.
func Int16(l string, v int16) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int32 returns LV.
func Int32(l string, v int32) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int64 returns LV.
func Int64(l string, v int64) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatInt(v, 10),
	}
}

// Uint returns LV.
func Uint(l string, v uint) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint8 returns LV.
func Uint8(l string, v uint8) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint16 returns LV.
func Uint16(l string, v uint16) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint32 returns LV.
func Uint32(l string, v uint32) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint64 returns LV.
func Uint64(l string, v uint64) *LV {
	return &LV{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Float32 returns LV of float32.
func Float32(l string, v float32) *LV {
	return &LV{
		Label: l,
		Value: fmt.Sprint(v),
	}
}

// Float63 returns LV of float64.
func Float63(l string, v float64) *LV {
	return &LV{
		Label: l,
		Value: fmt.Sprint(v),
	}
}

// Time returns LV of time.
func Time(l string, v time.Time) *LV {
	return &LV{
		Label: l,
		Value: v.Format(time.RFC3339Nano),
	}
}

// UTCTime returns LV of UTC time.
func UTCTime(l string, v time.Time) *LV {
	return Time(l, v.UTC())
}

// DefaultStackDepth is the depth used when call Stack.
var DefaultStackDepth = 32

// Stack returns LV of stack trace.
// depth is determined by DefaultStackDepth.
func Stack(l string) *LV {
	return &LV{
		Label: l,
		Value: stack.Callers(DefaultStackDepth, 1).String(),
	}
}

// StackN returns LV of stack trace for N layers.
func StackN(l string, depth int) *LV {
	return &LV{
		Label: l,
		Value: stack.Callers(depth, 1).String(),
	}
}
