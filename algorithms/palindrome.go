package algorithms

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	res := []rune{}
	for _, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			res = append(res, unicode.ToLower(r))
		}
	}
	left := 0
	right := len(res) - 1

	for left < right {
		if res[left] != res[right] {
			return false
		}
		left++
		right--
	}
	return true
}