package reviews

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= current[1] {
			if intervals[i][1] > current[1] {
				current[1] = intervals[i][1]
			}
		} else {
			result = append(result, current)
			current = intervals[i]
		}
	}
	result = append(result, current)
	return result
}