package algorithms

import "sort"

func GroupAnagram(strs []string) [][]string {
	m := make(map[string][]string)

	res := [][]string{}
	for _, str := range strs {
		char := []rune(str)
		sort.Slice(char, func(i, j int) bool {
			return char[i] < char[j]
		})
		s := string(char)
		m[s] = append(m[s], str)
	}

	for _, group := range m {
		res = append(res, group)
	}
	return res
}