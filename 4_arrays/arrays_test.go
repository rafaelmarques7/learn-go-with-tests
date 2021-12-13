package arrays

import (
	"reflect"
	"testing"
)

func TestArrays(t *testing.T) {
	assertEqual := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("Error - got %d, want %d", got, want)
		}
	}

	checkSumOfSlices := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error - got %d, want %d", got, want)
		}
	}

	t.Run("SumOfArray", func(t *testing.T) {
		// arrays have a fixed length that must be declared
		arr := [3]int{1, 2, 3}
		// alternatively, you can declare it this way
		// arr := [...]int{1, 2, 3}

		got := SumOfArray(arr)
		want := 6
		assertEqual(t, got, want)
	})

	t.Run("SumOfSlices", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		got := SumOfSlice(slice)
		want := 10
		assertEqual(t, got, want)
	})

	t.Run("SumEachSlice", func(t *testing.T) {
		got := SumEachSlice([]int{1, 2}, []int{10, 20})
		want := []int{3, 30}

		// The code below is not valid, because "(slice can only be compared to nil)"
		// if got != want {
		// 	t.Errorf("Error - got %v, want %v", got, want)
		// }

		// An alternative is to use reflect.DeepValue to compare two variables of any type!
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error - got %v, want %v", got, want)
		}
	})

	t.Run("SumAllTails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{10, 20, 40})
		want := []int{5, 60}
		checkSumOfSlices(t, got, want)
	})

	t.Run("SumAllTails handles empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{10, 20, 40})
		want := []int{0, 60}
		checkSumOfSlices(t, got, want)
	})
}
