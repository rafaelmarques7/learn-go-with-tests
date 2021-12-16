package bdt

import (
	"fmt"
	"math"
)

func ConvertToRoman(num int) string {
	main_symbol_values := []int{1000, 100, 10, 1}
	main_symbol_figures := []string{"M", "C", "X", "I"}

	// aux_symbol_values := []int{500, 50, 5}
	aux_symbol_figures := []string{"D", "L", "V"}

	roman_value := ""
	remaining := num

	fmt.Printf("inside convert\n")
	for index, val := range main_symbol_values {
		// check if num is divisible by val
		floor := int(math.Floor(float64(remaining) / float64(val)))

		if floor >= 1 && floor <= 3 {
			// if it is, take 1, 2 or 3 of that symbol at most
			for i := 0; i < floor; i++ {
				roman_value += main_symbol_figures[index]
			}
		} else if floor == 4 {
			// here we need to do the subtraction thing
			roman_value += main_symbol_figures[index]
			roman_value += aux_symbol_figures[index-1]
		} else if floor == 9 {
			// here we need to do the subtraction thing
			roman_value += main_symbol_figures[index]
			roman_value += main_symbol_figures[index-1]
		} else if floor >= 6 && floor <= 8 {
			roman_value += aux_symbol_figures[index-1]
			for i := 0; i < floor-5; i++ {
				roman_value += main_symbol_figures[index]
			}
		}

		// subtract the added ammount from remaining
		// and continue the iteration
		remaining = remaining - floor*val
	}

	return roman_value
}
