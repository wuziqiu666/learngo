package arraysandslices

// Sum returns the sum of the numbers
func Sum(numbers []int) int {
	add := func(acc, x int) int {
		return acc + x
	}
	return Reduce[int](numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTails := func(acc, x []int) []int {
		if len(x) == 0 {
			acc = append(acc, 0)
		} else {
			x := x[1:]
			acc = append(acc, Sum(x))
		}
		return acc
	}
	return Reduce[[]int](numbersToSum, sumTails, []int{})
}

func Reduce[T any](collection []T, accumulator func(T, T) T, initialValue T) T {
	result := initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
