package algorithms

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count := make(map[rune]int)
	
		for _, r := range s1 {
			count[r]++
		}

		for _, r := range s2 {
			count[r]--
			if count[r] < 0 {
				return false
			}
		}
	return true
}