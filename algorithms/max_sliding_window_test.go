package algorithms

import "testing"

func TestMaxSlidingWindow(t *testing.T) {
    got := MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
    want := []int{3, 3, 5, 5, 6, 7}
    
    if len(got) != len(want) {
        t.Fatalf("length mismatch: got %d, want %d", len(got), len(want))
    }
    
    for i := range got {
        if got[i] != want[i] {
            t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
        }
    }
}