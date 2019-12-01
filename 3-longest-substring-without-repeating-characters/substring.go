package longestsubstring

import (
	"sort"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	substrings := split(s)
	if len(substrings) == 0 {
		return 0
	}

	sort.Sort(byLen(substrings))

	return len(substrings[0])
}

type byLen []string

func (a byLen) Len() int {
	return len(a)
}

func (a byLen) Less(i, j int) bool {
	return len(a[i]) > len(a[j])
}

func (a byLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func split(s string) []string {
	var substrings []string

	length := len(s)
	if length == 1 {
		substrings = append(substrings, s)
	}
	for i := 0; i < length; i++ {
		var cur string

		seen := make(map[string]struct{})

		for j := i; j < length; j++ {
			e := string(s[j])

			if _, exist := seen[e]; !exist {
				cur += e
				seen[e] = struct{}{}
			} else {
				break
			}
		}

		substrings = append(substrings, cur)
	}

	return substrings
}
