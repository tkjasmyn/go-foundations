package algorithms

func ValidPalindrome(s string) bool {
	left := 0
	right := len(s)-1
	
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return IsPalindrome(s, left+1, right) || IsPalindrome(s, left, right-1)
		}
	}
	return true
}

func IsPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}