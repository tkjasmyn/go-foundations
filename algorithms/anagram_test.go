package algorithms

import "testing"

func TestIsAnagram(t *testing.T) {
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
		got := IsAnagram(tt.s1, tt.s2)
		if got != tt.want {
			t.Errorf("IsAnagram(%q, %q) = %v want %v", tt.s1, tt.s2, got, tt.want)
		}
	}
}