package main

import "testing"

// testing function must start with the word "Test"
// note that it must be capitalized
// the function must take exactly one argument `t *testing.T`
func TestSayHello(t *testing.T) {
	got := sayHello()
	want := "Hello World!"

	if got != want {
		t.Errorf("Unexpected result. Got %q, want %q", got, want)
	}
}

func TestSayHelloToAnyone(t *testing.T) {
	got := sayHelloToAnyone("Raf")
	want := "Hello, Raf!"

	if got != want {
		t.Errorf("Unexpected result. Got %q, want %q", got, want)
	}
}