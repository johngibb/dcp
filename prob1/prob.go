package prob1

func ContainsSumPairNaive(vals []int, sum int) bool {
	for i, a := range vals {
		for j, b := range vals {
			if i != j && a+b == sum {
				return true
			}
		}
	}
	return false
}

func ContainsSumPairMap(vals []int, sum int) bool {
	wanted := make(map[int]int)
	for i, v := range vals {
		wanted[sum-v] = i
	}
	for i, v := range vals {
		if j, ok := wanted[v]; i != j && ok {
			return true
		}
	}
	return false
}
