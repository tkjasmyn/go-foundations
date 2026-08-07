package reviews

import "testing"

func TestTwoSumII(t *testing.T) {
    got := TwoSumII([]int{2, 7, 11, 15}, 9)
    want := []int{1, 2}
    
    if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
        t.Errorf("got %v, want %v", got, want)
    }
}