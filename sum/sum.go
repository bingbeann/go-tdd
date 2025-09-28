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
		var sum int
		if len(nums) > 0 {
			tail := nums[1:]
			sum = Sum(tail)
		}

		sums[i] = sum
	}

	return sums
}
