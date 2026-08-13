package reviews

import "testing"

func TestMaxArea(t *testing.T) {
    tests := []struct {
        height []int
        want   int
    }{
        {[]int{1,8,6,2,5,4,8,3,7}, 49},
        {[]int{1,1}, 1},
        {[]int{4,3,2,1,4}, 16},
    }
    for _, tt := range tests {
        got := MaxArea(tt.height)
        if got != tt.want {
            t.Errorf("MaxArea(%v) = %d, want %d", tt.height, got, tt.want)
        }
    }
}