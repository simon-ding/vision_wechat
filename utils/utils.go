package utils

import (
	"unicode"
)

func IsChineseChar(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}
