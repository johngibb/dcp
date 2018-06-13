package prob004

// Given an array of integers, find the first missing positive integer in linear
// time and constant space. In other words, find the lowest positive integer
// that does not exist in the array. The array can contain duplicates and
// negative numbers as well.
//
// For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0]
// should give 3.
//
// You can modify the input array in-place.

func FindMissingPositiveIntegerNaive(vals []int) int {
outer:
	for i := 1; i <= len(vals); i++ {
		for _, v := range vals {
			if v == i {
				continue outer
			}
		}
		return i
	}
	return len(vals) + 1
}

func FindMissingPositiveIntegerEfficient(vals []int) int {
	// Since we can mutate the array, we can utilize a clever hack: set the sign
	// of the element at index x based on the existence of x+1 in the array. For
	// example, go through the array, and if we see a 1, flip `vals[1]` to
	// negative. Because there can be negatives, we need to a first pass to set
	// them to zero so they'll be ignored.

	// Set negatives to zero.
	for i, v := range vals {
		if v < 0 {
			vals[i] = 0
		}
	}

	// Now, set the sign on elements that we find in the positive sequential
	// series.
	for _, v := range vals {
		// The value might have been flipped to negative by a previous pass.
		v = abs(v)

		// If v is within the desired positive range, flip the sign of the
		// corresponding element.
		if v > 0 && v <= len(vals) {
			vals[v-1] = -abs(vals[v-1]) // -1 because it's 1,2,3…, not 0,1,2…
		}
	}

	// Now find the first positive index.
	for i, v := range vals {
		if v >= 0 {
			return i + 1
		}
	}
	return len(vals) + 1
}

// This is an adaptation of Dan's solution.
func FindMissingPositiveIntegerAlternate(a []int) int {
	for i := 0; i < len(a); i++ {
		v := a[i]
		if v <= 0 || v > len(a) {
			continue
		}
		j := v - 1 // desired index
		if a[j] == v {
			continue
		}
		a[i], a[j] = a[j], a[i]
		i--
	}
	for i, v := range a {
		if v != i+1 {
			return i + 1
		}
	}
	return len(a) + 1
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
