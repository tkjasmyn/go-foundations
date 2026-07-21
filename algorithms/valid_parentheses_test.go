package algorithms

import "testing"

func TestValidParentheses(t *testing.T) {
    tests := []struct {
        input string
        want  bool
    }{
        {"()", true},
        {"()[]{}", true},
        {"(]", false},
        {"([)]", false},
        {"{[]}", true},
        {"", true},
        {"(", false},
        {"((", false},
    }

    for _, tt := range tests {
        got := ValidParentheses(tt.input)
        if got != tt.want {
            t.Errorf("ValidParentheses(%q) = %v, want %v", tt.input, got, tt.want)
        }
    }
}