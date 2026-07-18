package algorithms

func TwoSumReview(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		comp := target - num
		if _, ok := seen[comp]; ok {
			return []int{seen[comp], i}
		}
		seen[num] = i
	}
	return nil
}

func IsAnagramReview(s1, s2 string) bool {
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