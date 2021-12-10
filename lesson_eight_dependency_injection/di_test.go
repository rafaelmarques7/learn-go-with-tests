package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Raf")

	got := buffer.String()
	want := "Hello, Raf"

	if got != want {
		t.Errorf("unexpected error - got %q, want %q", got, want)
	}
}
