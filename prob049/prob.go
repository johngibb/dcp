package prob

// This problem was asked by Amazon.
//
// Given an array of numbers, find the maximum sum of any contiguous subarray of
// the array.
//
// For example, given the array [34, -50, 42, 14, -5, 86], the maximum sum would
// be 137, since we would take elements 42, 14, -5, and 86.
//
// Given the array [-5, -1, -8, -9], the maximum sum would be 0, since we would
// not take any elements.
//
// Do this in O(N) time.

func MaxContiguousSum(vals []int) int {
	var sum, start int
	// Find the last index for which all previous values sum to a positive
	// number.
	for i, v := range vals {
		if sum+v < 0 {
			start = i + 1
			sum = 0
			continue
		}
		sum += v
	}
	sum = 0
	end := len(vals) - 1
	for i := len(vals) - 1; i >= start; i-- {
		if sum+vals[i] < 0 {
			end = i - 1
			sum = 0
		}
		sum += vals[i]
	}

	result := 0
	for i := start; i <= end; i++ {
		result += vals[i]
	}

	return result
}

// MaxContiguousSumKadane is an implementation of Kadane's algorithm.
func MaxContiguousSumKadane(vals []int) int {
	maxHere := vals[0]
	maxSoFar := maxHere
	for _, v := range vals[1:] {
		maxHere = max(v, maxHere+v)
		maxSoFar = max(maxHere, maxSoFar)
	}
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
