package algorithms

import "testing"

func TestValidPalindrome(t *testing.T) {
    tests := []struct {
        input string
        want  bool
    }{
        {"aba", true},
        {"abca", true},
        {"abc", false},
        {"deeee", true},
        {"a", true},
        {"", true},
    }
    
    for _, tt := range tests {
        got := ValidPalindrome(tt.input)
        if got != tt.want {
            t.Errorf("ValidPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
        }
    }
}