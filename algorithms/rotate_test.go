package algorithms

import "testing"

func TestRotate(t *testing.T) {
    nums := []int{1, 2, 3, 4, 5, 6, 7}
    Rotate(nums, 3)
    want := []int{5, 6, 7, 1, 2, 3, 4}
    
    for i := range nums {
        if nums[i] != want[i] {
            t.Errorf("index %d: got %d, want %d", i, nums[i], want[i])
        }
    }
}