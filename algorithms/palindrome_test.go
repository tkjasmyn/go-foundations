package algorithms

import "testing"

func TestIsPalindrome(t *testing.T) {
	test := []struct {
		input string
		want bool
	} {
		{"racecar", true},
		{"hello", false},
		{"a", true},
		{"", true},
	}

	for _, tt := range test {
		got := isPalindrome(tt.input)
		if got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}