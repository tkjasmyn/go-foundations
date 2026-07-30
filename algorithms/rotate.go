package algorithms

func Rotate(nums []int, k int)  {
	if len(nums) == 0 {
		return
	}
	res := make([]int, len(nums))
	k = k % len(nums)

	for i := 0; i < len(nums); i++ {
		res[(i+k) % len(nums)] = nums[i]
	}
	copy(nums, res)
}