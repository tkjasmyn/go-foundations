package reviews

func MaxArea(height []int) int {
	left := 0
	right := len(height)-1
	var max int

	if len(height) < 2 {
		return -1	
	}

	for left < right {
		width := right - left
		h := min(height[left], height[right])
		water := width * h

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}

		if water > max {
			max = water
		}
	}
	return max
}