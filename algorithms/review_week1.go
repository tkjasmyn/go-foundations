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