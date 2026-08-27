package reviews

func LengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	seen := make(map[byte]bool)
	max := 0

	for right < len(s) {
		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}

		seen[s[right]] = true
		win := right - left + 1
		if win > max {
			max = win
		}
		right++
	}
	return max
}