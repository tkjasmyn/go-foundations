package algorithms

func ProductExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	leftproduct := 1
	for i := 0; i < len(nums); i++ {
		output[i] = leftproduct
		leftproduct = leftproduct * nums[i]
	}
	rightproduct := 1
	for i := len(nums)-1; i >= 0; i-- {
		output[i] *= rightproduct
		rightproduct *= nums[i]
	}
	return output
}