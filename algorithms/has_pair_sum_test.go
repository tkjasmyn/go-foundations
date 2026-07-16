package algorithms

import "testing"

func TestHasPairSum(t *testing.T) {
    tests := []struct {
        nums   []int
        target int
        want   bool
    }{
        {[]int{2, 7, 11, 15}, 9, true},
        {[]int{1, 2, 3}, 10, false},
        {[]int{3, 3}, 6, true},
    }

    for _, tt := range tests {
        got := HasPairSum(tt.nums, tt.target)
        if got != tt.want {
            t.Errorf("HasPairSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
        }
    }
}