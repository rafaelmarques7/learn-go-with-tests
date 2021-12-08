package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("Unexpected summ, got %d, expected %d", sum, expected)
	}
}

// This example is actually evaluated as part of the tests (run `go test -v` to see this)
// and it will fail case the function output does not match the `Output` as per the comment
// Note: function name must start with "Example", and it must be capitalized
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
