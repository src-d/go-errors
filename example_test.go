package errors_test

import (
	"fmt"
	"io"

	"srcd.works/errors.v0"
)

func ExampleNew() {
	var ErrExample = errors.NewKind("example")

	err := ErrExample.New()
	if ErrExample.Is(err) {
		fmt.Println(err)
	}

	// Output: example
}

func ExampleNew_format() {
	var ErrMaxLimitReached = errors.NewKind("max. limit reached: %d")

	err := ErrMaxLimitReached.New(42)
	if ErrMaxLimitReached.Is(err) {
		fmt.Println(err)
	}

	// Output: max. limit reached: 42
}

func ExampleKind_Wrap() {
	var ErrNetworking = errors.NewKind("network error")

	err := ErrNetworking.Wrap(io.EOF)
	if ErrNetworking.Is(err) {
		fmt.Println(err)
	}

	// Output: network error: EOF
}

func ExampleKind_Wrap_nested() {
	var ErrNetworking = errors.NewKind("network error")
	var ErrReading = errors.NewKind("reading error")

	err3 := io.EOF
	err2 := ErrReading.Wrap(err3)
	err1 := ErrNetworking.Wrap(err2)
	if ErrReading.Is(err1) {
		fmt.Println(err1)
	}

	// Output: network error: reading error: EOF
}

func ExampleAny() {
	var ErrNetworking = errors.NewKind("network error")
	var ErrReading = errors.NewKind("reading error")

	err := ErrNetworking.New()
	if errors.Any(err, ErrReading, ErrNetworking) {
		fmt.Println(err)
	}

	// Output: network error
}
