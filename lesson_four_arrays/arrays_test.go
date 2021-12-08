package arrays

import "testing"

func TestSumOfArray(t *testing.T) {
	// arrays have a fixed length that must be declared
	arr := [3]int{1, 2, 3}
	// alternatively, you can declare it this way
	// arr := [...]int{1, 2, 3}
	got := SumOfArray(arr)
	want := 6

	if got != want {
		t.Errorf("Error sum - got %d, want %d, given input argument %v", got, want, arr)
	}
}

func TestSumOfSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	got := SumOfSlice(slice)
	want := 6
	if got != want {
		t.Errorf("Error sum - got %d, want %d, given input argument %v", got, want, slice)
	}
}
