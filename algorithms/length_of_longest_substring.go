package algorithms

func LengthOfLongestSubstring(s string) int {
	count := 0
	left := 0
	right := 0
	seen := make(map[byte]bool)

	for right < len(s) {
		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}
		seen[s[right]] = true
		win := right - left + 1
		if win > count {
			count = win
		}
		right++
	}
	return count
}