package reviews

func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums)-1

	for left <= right {
		mid := (right + left) / 2
		if target == nums[mid] {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		}
		
		if target < nums[mid] {
			right = mid - 1
		}
	}
	return -1
}