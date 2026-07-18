package algorithms

import "testing"

func TestTwoSumReview(t *testing.T) {
	tests := []struct {
		nums []int
		target int
		want []int
	} {
		{[]int{2, 7, 11, 13}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		got := TwoSumReview(tt.nums, tt.target)
		if len(got) != 2 || got[0] != tt.want[0] || got[1] != tt.want[1] {
			t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}