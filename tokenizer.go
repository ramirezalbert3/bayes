package tokenizer

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
