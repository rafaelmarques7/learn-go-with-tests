package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertEqual := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Error - got %q, want %q", got, want)
		}
	}

	t.Run("Repeat function takes a string and repeats it 5 times by default", func(t *testing.T) {
		got := Repeat("a", -1)
		want := "aaaaa"
		assertEqual(t, got, want)
	})

	t.Run("Repeat function allows you to specify how many times to repeat the string", func(t *testing.T) {
		got := Repeat("cool", 3)
		want := "coolcoolcool"
		assertEqual(t, got, want)

	})

}

// This is a benchmark test
func BenchmarkRepeat(b *testing.B) {
	// Note: b.N is a constant determined by the benchmark library
	// it chooses this value based on a "good" value to test the function
	// and frees you from having to make that decision yourself
	for i := 0; i < b.N; i++ {
		Repeat("aString", -1)
	}
}

func ExampleRepeat() {
	repeated := Repeat("Penny?", 3)
	fmt.Println(repeated)
	// Output: Penny?Penny?Penny?
}
