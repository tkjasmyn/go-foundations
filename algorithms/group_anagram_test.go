package algorithms

import "testing"

func TestGroupAnagram(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    got := GroupAnagram(input)
    
    if len(got) != 3 {
        t.Errorf("Expected 3 groups, got %d", len(got))
    }
}