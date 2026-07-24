package algorithms

import "testing"

func TestSearch(t *testing.T) {
    tests := []struct {
        nums   []int
        target int
        want   int
    }{
        {[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
        {[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
        {[]int{1, 2, 3, 4, 5}, 2, 1},
        {[]int{5}, 5, 0},
        {[]int{5}, -5, -1},
    }

    for _, tt := range tests {
        got := Search(tt.nums, tt.target)
        if got != tt.want {
            t.Errorf("Search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
        }
    }
}