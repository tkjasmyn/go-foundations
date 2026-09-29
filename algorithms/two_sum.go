package algorithms

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]

		if _, ok := seen[comp]; ok {
			return []int{seen[comp], i}
		}
		seen[nums[i]] = i
	}
	return nil
}