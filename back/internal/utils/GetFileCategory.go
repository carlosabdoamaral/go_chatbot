package utils

import (
	"fmt"
)

func GetFileCategory(s string) string {
	s = _removeExtension(s)
	s = _removePath(s)
	return s
}

func _removeExtension(s string) string {
	res := ""
	indexes := []int{}

	for i, char := range s {
		if string(char) != "." {
			indexes = append(indexes, i)
		} else {
			break
		}
	}

	for _, index := range indexes {
		res = fmt.Sprintf("%s%s", res, string(s[index]))
	}

	return res
}

func _removePath(s string) string {
	var (
		res           = ""
		lastDashIndex = 0
	)

	for i, char := range s {
		if string(char) == "/" {
			lastDashIndex = i
		}
	}

	for k := lastDashIndex + 1; k < len(s); k++ {
		res = fmt.Sprintf("%s%s", res, string(s[k]))
	}

	return res
}
