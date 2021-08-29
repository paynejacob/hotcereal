package graph

import (
	"strings"
)

func Tokenize(in string) []string {
	if len(in) == 0 {
		return []string{}
	}

	tokens := strings.Split(in, " ")

	for i, _ := range tokens {
		tokens[i] = strings.ToLower(tokens[i])
	}

	return tokens
}
