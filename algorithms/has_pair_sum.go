package algorithms

func HasPairSum(nums []int, target int) bool {
	seen := make(map[int]int)
	for _, num := range nums {
		comp := target-num

		if _, ok := seen[comp]; ok {
			return true
		}
		seen[num] = num
	}
	return false
}