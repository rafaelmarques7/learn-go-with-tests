package arrays

// Note:
// arrays have fixed length
// slices have variable length
// the array size is encoded in its type -> an array of different size is a different type!

func SumOfArray(array [3]int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func SumOfSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}
