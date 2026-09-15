package reviews

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 10, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, c := range cases {
		Rotate(c.nums, c.k)
		if !reflect.DeepEqual(c.nums, c.want) {
			t.Errorf("Rotate(_, %d) = %v, want %v", c.k, c.nums, c.want)
		}
	}
}