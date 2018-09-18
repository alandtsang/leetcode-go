package ispalindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) <= 1 { // 空串和单个字符
		return true
	}

	str := strings.ToLower(s)
	var i, j int

	for i, j := 0, len(str)-1; i < j; {
		if unicode.IsLetter(rune(str[i])) == false &&
			unicode.IsDigit(rune(str[i])) == false {
			i++
			continue
		}
		if unicode.IsLetter(rune(str[j])) == false &&
			unicode.IsDigit(rune(str[j])) == false {
			j--
			continue
		}
		if str[i] == str[j] {
			i++
			j--
			continue
		} else {
			return false
		}
	}

	if i == j { // eg: "a."
		return true
	} else {
		return false
	}
}
