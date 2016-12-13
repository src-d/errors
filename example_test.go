package errors_test

import (
	"fmt"
	"io"

	"srcd.works/errors.v0"
)

func ExampleNew() {
	var ErrExample = errors.New("example")

	err := ErrExample.New()
	if ErrExample.Is(err) {
		fmt.Println(err)
	}

	// Output: example
}

func ExampleNewFormat() {
	var ErrMaxLimitReached = errors.New("max. limit reached: %d")

	err := ErrMaxLimitReached.New(42)
	if ErrMaxLimitReached.Is(err) {
		fmt.Println(err)
	}

	// Output: max. limit reached: 42
}

func ExampleWrap() {
	var ErrNetworking = errors.New("network error")

	err := ErrNetworking.Wrap(io.EOF)
	if ErrNetworking.Is(err) {
		fmt.Println(err)
	}

	// Output: network error: EOF
}

func ExampleNestedWrap() {
	var ErrNetworking = errors.New("network error")
	var ErrReading = errors.New("reading error")

	err3 := io.EOF
	err2 := ErrReading.Wrap(err3)
	err1 := ErrNetworking.Wrap(err2)
	if ErrReading.Is(err1) {
		fmt.Println(err1)
	}

	// Output: network error: reading error: EOF
}

func ExampleAlternativeMultiple() {
	var ErrNetworking = errors.New("network error")
	var ErrReading = errors.New("reading error")

	err := ErrNetworking.New()
	if errors.Is(err, ErrReading, ErrNetworking) {
		fmt.Println(err)
	}

	// Output: network error
}