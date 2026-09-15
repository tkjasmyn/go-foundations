package reviews

func Rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	if k > len(nums) {
		k = k % len(nums)
	}

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	left := 0
	right := len(nums)-1

	for left <= right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}