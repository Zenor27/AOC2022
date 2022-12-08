package utils

func RemoveFromList(l []int, valueToRemove int) []int {
	for i, v := range l {
		if v == valueToRemove {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func GetMaxFromList(l []int) int {
	max := 0
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return max
}
