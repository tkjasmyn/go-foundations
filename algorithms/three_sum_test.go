package algorithms

import "testing"

func TestThreeSum(t *testing.T) {
	got := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	want := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %d, want %d", len(got), len(want))
	}

	for i := 0; i < len(got); i++ {
		if len(got[i]) != len(want[i]) {
			t.Fatalf("triplet %d length mismatch", i)
		}
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				t.Errorf("triplet %d, index %d: got %d, want %d", i, j, got[i][j], want[i][j])
			}
		}
	}
}