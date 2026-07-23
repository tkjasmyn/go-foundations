package algorithms

func MaxSlidingWindow(nums []int, k int) []int {
	res := []int{}

	for i := 0; i <= len(nums)-k; i++ {
		glass := nums[i : i+k]
		max := max(glass)
		res = append(res, max)
	}
	return res
}

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}