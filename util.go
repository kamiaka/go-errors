package errors

import (
	"strings"
	"time"
)

var escaper = strings.NewReplacer("\t", "\\t", "\n", "\\n", "\\", "\\\\")

func escape(s string) string {
	return escaper.Replace(s)
}

var digits = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func appendHexBytes(buf []byte, v []byte) []byte {
	buf = append(buf, "0x"...)
	for _, b := range v {
		buf = append(buf, digits[b/16])
		buf = append(buf, digits[b%16])
	}
	return buf
}

func appendHexByte(buf []byte, b byte) []byte {
	buf = append(buf, "0x"...)
	buf = append(buf, digits[b/16])
	buf = append(buf, digits[b%16])
	return buf
}

func appendUTCTime(buf []byte, t time.Time) []byte {
	t = t.UTC()
	tmp := []byte("0000-00-00T00:00:00.000000Z")
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	itoa(tmp[:4], year, 4)
	itoa(tmp[5:7], int(month), 2)
	itoa(tmp[8:10], day, 2)
	itoa(tmp[11:13], hour, 2)
	itoa(tmp[14:16], min, 2)
	itoa(tmp[17:19], sec, 2)
	itoa(tmp[20:26], t.Nanosecond()/1e3, 6)
	return append(buf, tmp...)
}

// Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
// Copied from https://github.com/golang/go/blob/go1.8.1/src/log/log.go#L75-L90
// and modified for ltsvlog.
// It is user's responsibility to pass buf which len(buf) >= wid
func itoa(buf []byte, i int, wid int) {
	// Assemble decimal in reverse order.
	bp := wid - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		buf[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	buf[bp] = byte('0' + i)
}
