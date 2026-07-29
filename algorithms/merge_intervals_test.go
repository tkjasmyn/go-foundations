package algorithms

import "testing"

func TestMerge(t *testing.T) {
    got := Merge([][]int{{1,3},{2,6},{8,10},{15,18}})
    want := [][]int{{1,6},{8,10},{15,18}}
    
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %v, want %v", len(got), len(want))
	}

	for i := 0; i < len(got); i++ {
		if len(got[i]) != len(want[i]) {
			t.Fatalf("interval %d length mismatch", i)
		}
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				t.Errorf("interval %d, index %d: got %d, want %d", i, j, got[i][j], want[i][j])
			}
		}
	}
}