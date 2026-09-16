package reviews

import "sort"

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	res := [][]string{}

	for _, word := range strs {
		char := []rune(word)

		sort.Slice(char, func(i, j int) bool {
			return char[i] < char[j]
		})

		str := string(char)

		groups[str] = append(groups[str], word)
	}

	for _, group := range groups {
		res = append(res, group)
	}
	return res
}