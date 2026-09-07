package reviews

import "sort"

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	if len(strs) == 0 {
		return [][]string{}
	}

	res := [][]string{}

	for _, word := range strs {
		chars := []rune(word)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		str := string(chars)
		groups[str] = append(groups[str], word)
	}

	for _, group := range groups {
		res = append(res, group)
	}
	return res
}