package reviews

func MaxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxArea := 0

    for left < right {
        width := right - left
        h := min(height[left], height[right])

        area := width * h
        if area > maxArea {
            maxArea = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return maxArea
}