package reviews

import (
	"sort"
	"testing"
)

func sortSlices(slices [][]string) [][]string {
	for _, s := range slices {
		sort.Strings(s)
	}
	sort.Slice(slices, func(i, j int) bool {
		if len(slices[i]) == 0 || len(slices[j]) == 0 {
			return len(slices[i]) < len(slices[j])
		}
		return slices[i][0] < slices[j][0]
	})
	return slices
}
func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
		{
			[]string{},
			[][]string{},
		},
	}

	for _, tt := range tests {
		got := GroupAnagrams(tt.input)
		got = sortSlices(got)
		expected := sortSlices(tt.expected)

		if len(got) != len(expected) {
			t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.input, got, expected)
			continue
		}

		for i := 0; i < len(got); i++ {
			if len(got[i]) != len(expected[i]) {
				t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.input, got, expected)
				break
			}
			for j := 0; j < len(got[i]); j++ {
				if got[i][j] != expected[i][j] {
					t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.input, got, expected)
					break
				}
			}
		}
	}
}