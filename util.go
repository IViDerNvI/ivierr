package ivierr

func computePrefixTable(pattern string) []int {
	m := len(pattern)
	prefixTable := make([]int, m)
	j := 0

	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = prefixTable[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		prefixTable[i] = j
	}

	return prefixTable
}

func kmpSearch(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	prefixTable := computePrefixTable(pattern)
	result := []int{}
	j := 0

	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = prefixTable[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			result = append(result, i-m+1)
			j = prefixTable[j-1]
		}
	}

	return result
}
