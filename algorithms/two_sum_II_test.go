package algorithms

import "testing"

func TestTwoSumII(t *testing.T) {
    tests := []struct {
        nums   []int
        target int
        want   []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{1, 2}},
        {[]int{2, 3, 4}, 6, []int{1, 3}},
        {[]int{-1, 0}, -1, []int{1, 2}},
    }

    for _, tt := range tests {
        got := TwoSumII(tt.nums, tt.target)
        if len(got) != len(tt.want) || got[0] != tt.want[0] || got[1] != tt.want[1] {
            t.Errorf("TwoSumII(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
        }
    }
}