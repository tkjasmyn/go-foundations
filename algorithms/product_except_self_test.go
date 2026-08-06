package algorithms

import "testing"

func TestProductExceptSelf(t *testing.T) {
    got := ProductExceptSelf([]int{1, 2, 3, 4})
    want := []int{24, 12, 8, 6}
    
    if len(got) != len(want) {
        t.Fatalf("length mismatch")
    }
    for i := range got {
        if got[i] != want[i] {
            t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
        }
    }
}