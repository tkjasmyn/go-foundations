package algorithms

import "sort"

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		left := i+1
		right := len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]
			target := -nums[i]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
			if sum > target {
				right--
			}
			if sum < target {
				left++
			}
		}
	}
	return res
}