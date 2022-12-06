package utils

import "regexp"

func SplitByEmptyLine(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)
}

func HasOnlyUniqChars(str string) bool {
	m := make(map[rune]bool)
	for _, c := range str {
		_, ok := m[c]
		if ok {
			return false
		}

		m[c] = true
	}

	return true
}
