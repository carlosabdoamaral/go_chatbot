package utils

import (
	"fmt"
	"strings"
)

// ARRUMAR!
func CapitalizeWord(s string) string {
	var (
		res = ""
	)

	for i, char := range s {
		c := string(char)

		if i == 0 {
			res = fmt.Sprintf("%s%s", res, strings.ToUpper(c))
		} else {
			res = fmt.Sprintf("%s%s", res, strings.ToLower(c))
		}
	}

	return res
}
