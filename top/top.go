package top

import (
	"sort"
	"strings"
)

type TopResult struct {
	Word  string
	Count int
}

func GetTopResults(text string, n int) ([]TopResult, error) {
	m := make(map[string]int)

	words := strings.Split(text, " ")

	for _, w := range words {
		lowercaseWord := strings.ToLower(w)
		if _, ok := m[lowercaseWord]; ok {
			m[lowercaseWord] = m[lowercaseWord] + 1
		} else {
			m[lowercaseWord] = 1
		}
	}

	results := make([]TopResult, len(m))
	i := 0
	for k, v := range m {
		results[i] = TopResult{
			Word:  k,
			Count: v,
		}

		i = i + 1
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Count > results[j].Count
	})

	if len(results) > n {
		return results[:n], nil
	}

	return results, nil
}
