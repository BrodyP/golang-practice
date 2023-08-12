package array_slices

func SumAll(numbersToSum ...[]int) []int {
	result := []int{}
	for _, numberArray := range numbersToSum {
		result = append(result, Sum(numberArray))
	}

	return result
}
