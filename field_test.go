package errors

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestWithField_Format(t *testing.T) {
	cases := []struct {
		err  error
		want []byte
	}{
		{
			err: WithFields(fmt.Errorf("Some error")).String("a", "value of a\nnext line").Int("b", 42),
			want: []byte("Some error\n\ta:value of a\\nnext line	b:42"),
		},
	}
	for i, tc := range cases {
		buf := new(bytes.Buffer)
		_, err := fmt.Fprintf(buf, "%+v", tc.err)
		if err != nil {
			t.Errorf("#%d Format() occured error: %s", i, err)
		}
		got := buf.Bytes()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("#%d *withField.Format() \ngot %s\nwant %s", i, got, tc.want)
		}
	}
}
