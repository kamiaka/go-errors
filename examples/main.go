package main

import (
	errs "errors"
	"fmt"

	errors "github.com/kamiaka/go-errors"
	"github.com/kamiaka/go-errors/lv"
)

var internalError = errors.NewType("internal error", errors.Level("warn"))
var runtimeError = errors.NewType("runtime error", errors.Level("info"))

var mustInput = internalError.New("must input name")

func find(name string) (string, error) {
	if name == "" {
		return "", mustInput
	}
	if name == "notExists" {
		return "", runtimeError.Errorf("%s is not find", name)
	}
	if name == "err" {
		return "", internalError.Errorf("%s is reserved", name)
	}
	return "ok", nil
}

func test(name string) error {
	val, err := find(name)
	if err != nil {
		if errs.Is(err, runtimeError) {
			fmt.Println(err)
			return nil
		}
		if !errs.Is(err, mustInput) {
			return internalError.Errorf("wrapped error: %w", err).With(lv.String("name", name), lv.Stack("stack"))
		}
		return err
	}
	fmt.Printf("%s: %s\n", name, val)
	return nil
}

func test2(name string) error {
	return test(name)
}

func main() {
	cases := []string{"", "exists", "noExists", "err"}
	for _, tc := range cases {
		err := test2(tc)
		if err != nil {
			fmt.Printf("%%+v: %+v\n", err)
		}
	}
}
