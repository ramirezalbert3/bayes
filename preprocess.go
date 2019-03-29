package preprocess

import (
	"regexp"
	"strings"
	)

var re_punctuation = regexp.MustCompile("[^a-zA-Z0-9-_ ]+")

func Tokenize(s string) []string {
	s = re_punctuation.ReplaceAllLiteralString(s, "")
	s = strings.ToLower(s)
	tokens := strings.Split(s, " ")
	return tokens
}

func CountVectorizer(texts []string) map[string]int {
	vocabulary := make(map[string]int)
	for _, text := range texts {
		tokens := Tokenize(text)
		for _, t := range tokens {
			if _, ok := vocabulary[t]; !ok {
				vocabulary[t] = 0
			}
			vocabulary[t] += 1
		}
	}
	return vocabulary
}