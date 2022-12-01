package utils

func RemoveFromList(l []int, valueToRemove int) []int {
	for i, v := range l {
		if v == valueToRemove {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
