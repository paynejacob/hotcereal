package graph

import (
	"regexp"
	"strings"
)

var tokenRexp = regexp.MustCompile("[^a-zA-Z0-9]+")

func Tokenize(in string) []string {
	if len(in) == 0 {
		return []string{}
	}

	tokens := strings.Split(in, " ")

	for i, _ := range tokens {
		tokens[i] = tokenRexp.ReplaceAllString(tokens[i], "")
		tokens[i] = strings.ToLower(tokens[i])
	}

	return tokens
}
