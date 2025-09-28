package sum

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAllTails(numbers ...[]int) []int {
	sums := make([]int, len(numbers))
	for i, nums := range numbers {
		tail := nums[1:]
		sums[i] = Sum(tail)
	}

	return sums
}
