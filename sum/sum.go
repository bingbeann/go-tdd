package sum

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, len(numbers))
	for i, nums := range numbers {
		sums[i] = Sum(nums)
	}

	return sums
}
