package errors_test

import (
	"fmt"
	"io"

	"srcd.works/errors.v0"
)

func ExampleKind_New() {
	var ErrExample = errors.NewKind("example")

	err := ErrExample.New()
	if ErrExample.Is(err) {
		fmt.Println(err)
	}

	// Output: example
}

func ExampleKind_New_pattern() {
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

func ExampleError_printf() {
	var ErrExample = errors.NewKind("example with stack trace")

	err := ErrExample.New()
	fmt.Printf("%+v\n", err)

	// Example Output:
	// example with stack trace
	//
	// srcd.works/errors%2v0_test.ExampleError_Format
	//         /home/mcuadros/workspace/go/src/srcd.works/errors.v0/example_test.go:60
	// testing.runExample
	//         /usr/lib/go/src/testing/example.go:114
	// testing.RunExamples
	//         /usr/lib/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /usr/lib/go/src/testing/testing.go:744
	// main.main
	//         github.com/pkg/errors/_test/_testmain.go:106
	// runtime.main
	//         /usr/lib/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /usr/lib/go/src/runtime/asm_amd64.s:2086

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
