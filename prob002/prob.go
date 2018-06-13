package prob2

// This problem was asked by Uber.

// Given an array of integers, return a new array such that each element at
// index i of the new array is the product of all the numbers in the original
// array except the one at i. Solve it without using division and in O(n) time.

// For example, if our input was [1, 2, 3, 4, 5], the expected output would be
// [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would
// be [2, 3, 6].

func ComputeNaive(input []int) []int {
	result := make([]int, len(input))
	for i := range input {
		prod := 1
		for j, v := range input {
			if i != j {
				prod *= v
			}
		}
		result[i] = prod
	}
	return result
}

func ComputeEfficient(input []int) []int {
	var (
		left   = make([]int, len(input))
		right  = make([]int, len(input))
		result = make([]int, len(input))
	)
	prod := 1
	for i, v := range input {
		left[i] = prod
		prod *= v
	}

	prod = 1
	for i := len(input) - 1; i >= 0; i-- {
		right[i] = prod
		prod *= input[i]
	}

	for i := range input {
		result[i] = left[i] * right[i]
	}
	return result
}

func ComputeEfficientOneArray(input []int) []int {
	result := make([]int, len(input))
	prod := 1
	for i, v := range input {
		result[i] = prod
		prod *= v
	}

	prod = 1
	for i := len(input) - 1; i >= 0; i-- {
		result[i] *= prod
		prod *= input[i]
	}
	return result
}

// [1, 2, 3, 4, 5]
