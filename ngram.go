package main

import (
	"strings"
)

func Ngram(line string) (ret [][]string) {
	tokens := cut(strings.TrimSpace(line))
	tokens = append(tokens, "<eos>")
	tokens = append([]string{"<sos>"}, tokens...)
	n := len(tokens)
	for i := 0; i < n; i++ {
		if strings.TrimSpace(tokens[i]) == "" {
			continue
		}
		for j := 1; j <= n; j++ {
			if i+j > n {
				break
			}
			if (j == 1) && (i == 0 || i == n-1) {
				continue
			}
			if strings.TrimSpace(tokens[i+j-1]) == "" {
				continue
			}
			ret = append(ret, tokens[i:i+j])
		}
	}
	return
}
