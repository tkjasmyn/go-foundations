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

func TestIsAnagramReview(t *testing.T)  {
	tests := []struct {
		s1, s2 string
		want bool
	} {
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"listen", "silent", true},
		{"a", "a", true},
		{"", "", true},
	}

	for _, tt := range tests {
		got := IsAnagramReview(tt.s1, tt.s2)
		if got != tt.want {
			t.Errorf("IsAnagram(%q, %q) = %v want %v", tt.s1, tt.s2, got, tt.want)
		}
	}
}