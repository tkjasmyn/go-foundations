package algorithms

import "testing"

func TestMaxArea(t *testing.T) {
    got := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
    want := 49
    if got != want {
        t.Errorf("MaxArea(...) = %d, want %d", got, want)
    }
}