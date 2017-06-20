package errors

import (
	"fmt"
	"testing"
)

func TestAppendCallers(t *testing.T) {
	buf := []byte{}
	buf = appendCallers(buf)
	fmt.Printf("buf: %s\n", buf)
}
