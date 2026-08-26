package reviews

func IsAnagram(s string, t string) bool {
	count := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, r := range s {
		count[r]++
	}

	for _, r := range t {
		count[r]--
		if count[r] < 0 {
			return false
		}
	}
	return true
}