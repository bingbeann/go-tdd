package sum

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, 0)
	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}

	return sums
}
