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


// package reviews

// import (
// 	"sort"
// 	"testing"
// )

// func TestGroupAnagrams(t *testing.T) {
// 	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
// 	got := GroupAnagrams(input)

// 	// sort each group and the outer slice so comparison isn't order-dependent
// 	for _, g := range got {
// 		sort.Strings(g)
// 	}
// 	sort.Slice(got, func(i, j int) bool { return got[i][0] < got[j][0] })

// 	want := [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}}

// 	if len(got) != len(want) {
// 		t.Fatalf("got %v groups, want %v", got, want)
// 	}
// 	for i := range want {
// 		if len(got[i]) != len(want[i]) {
// 			t.Errorf("group %d: got %v, want %v", i, got[i], want[i])
// 		}
// 		for j := range want[i] {
// 			if got[i][j] != want[i][j] {
// 				t.Errorf("group %d: got %v, want %v", i, got[i], want[i])
// 			}
// 		}
// 	}
// }