package algorithms

import "sort"

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})
	res := [][]int{}

	if len(intervals) == 0 {
		return res
	}

	current := intervals[0]

	for _, interval := range intervals[1:] {
		if interval[0] <= current[1] {
			if interval[1] > current[1] {
				current[1] = interval[1]
			}
		} else {
			res = append(res, current)
			current = interval
		}
	}
	res = append(res, current)
	return res
}