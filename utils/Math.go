package utils

func FindMax(l []int) int {
	if len(l) == 0 {
		return 0
	}

	max := l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}

	return max
}
