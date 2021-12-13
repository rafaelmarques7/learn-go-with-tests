package arrays

// Note:
// arrays have fixed length
// slices have variable length
// the array size is encoded in its type -> an array of different size is a different type!

func SumOfArray(array [3]int) int {
	// Note: this function takes an array of exactly size 3, and nothing else!
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func SumOfSlice(slice []int) int {
	// Note: this function takes a slice of any size!
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

// the three dotes (...) denote a variable number of arguments
// func SumEachSlice(slices ...[]int) []int {
// 	lenOfSlices := len(slices)
// 	sums := make([]int, lenOfSlices)

// 	for i, slice := range slices {
// 		sums[i] = SumOfSlice(slice)
// 	}
// 	return sums
// }

// Alternative implementation using append
func SumEachSlice(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		sums = append(sums, SumOfSlice(slice))
	}
	return sums
}

func SumAllTails(slices ...[]int) (sumTails []int) {
	for _, slice := range slices {
		if len(slice) > 0 {
			tail := slice[1:]
			sumTails = append(sumTails, SumOfSlice(tail))
		} else {
			sumTails = []int{0}
		}
	}
	return sumTails
}
