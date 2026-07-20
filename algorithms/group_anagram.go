package algorithms

import "sort"

func GroupAnagram(strs []string) [][]string {
	m := make(map[string][]string)
	res := [][]string{}

	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {return runes[i] < runes[j]})
		s := string(runes)
			m[s] = append(m[s], str)
	}

	for _, group := range m {
		res = append(res, group)
	}
	return res
}