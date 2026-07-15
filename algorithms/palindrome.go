package algorithms

func isPalindrome(s string) bool {
	runes := []rune(s)
	left := 0
	right := len(runes) - 1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}