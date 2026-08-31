package reviews

import "testing"

func TestMergeIntervals(t *testing.T)  {
	tests := []struct {
    	input    [][]int
    	expected [][]int
	}{
	    {[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
	    {[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	    {[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
	    {[][]int{}, [][]int{}},
	}
	
	for _, tt := range tests {
		got := MergeIntervals(tt.input)
		want := tt.expected

		if len(got) != len(want) {
			t.Errorf("MergeIntervals(%v) = %v, want %v", tt.input, got, want)
		}

		for i := 0; i < len(got); i++ {
			if got[i][0] != want[i][0] || got[i][1] != want[i][1] {
				t.Errorf("MergeIntervals(%v) = %v, want %v", tt.input, got, want)
			}
		}
	}
}