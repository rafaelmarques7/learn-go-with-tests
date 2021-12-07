package main

import "testing"

// testing function must start with the word "Test"
// note that it must be capitalized
// the function must take exactly one argument `t *testing.T`
func TestSayHelloToAnyone(t *testing.T) {
	// testing.TB is an interface that is satisfied by *testing.T and *testing.B
	assertCorrectHelloMessage := func(t testing.TB, got string, want string) {
		// t.Helper is useful to better identify where the error occurs
		// without this, the error would be indicated as happening inside this assert function
		// with the helper, the error is indicated as happening in the parent function that calls the assert
		t.Helper()
		if got != want {
			t.Errorf("Unexpected result. Got %q, want %q", got, want)
		}
	} 

	t.Run("Saying Hello while passing a persons name", func(t *testing.T) {
		got := SayHello("Raf", "")
		want := "Hello, Raf!"
		assertCorrectHelloMessage(t, got, want)
	})

	t.Run("Saying Hello without passing any arguments", func(t *testing.T) {
		got := SayHello("", "")
		want := "Hello, World!"
		assertCorrectHelloMessage(t, got, want)
	})

	t.Run("Saying Hello in Spanish", func(t *testing.T) {
		got := SayHello("Gemma", "Spanish")
		want := "Hola, Gemma!"
		assertCorrectHelloMessage(t, got, want)
	})

	t.Run("Saying Hello in French", func(t *testing.T) {
		got := SayHello("Emelie", "French")
		want := "Bonjour, Emelie!"
		assertCorrectHelloMessage(t, got, want)
	})
}