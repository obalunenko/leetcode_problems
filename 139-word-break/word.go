package wordbreak

func wordBreak(s string, wordDict []string) bool {
	if s == "" || len(wordDict) == 0 {
		return false
	}

	dict := make(map[string]bool, len(wordDict))

	for _, w := range wordDict {
		dict[w] = true
	}

	l := len(s)
	// remember indexes of next after breaks
	list := make(map[int]bool, l+1)
	list[0] = true

	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			cur := s[j:i]

			if dict[cur] && list[j] {
				list[i] = true

				break
			}
		}
	}

	return list[l]
}
