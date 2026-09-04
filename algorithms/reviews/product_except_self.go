package reviews

func ProductExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	for i := range output {
		output[i] = 1
	}

	leftProduct := 1
	for i := 0; i < len(nums); i++ {
		output[i] *= leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := len(nums)-1; i >= 0; i-- {
		output[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return output
}