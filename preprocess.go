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

// TODO: all this is very very naive
type CountVectorizer struct {
	Vocabulary []string
	WordCounts map[string]int
}

func (c *CountVectorizer) FitTransform(texts []string) (map[string]int, [][]int) {
	c.WordCounts = make(map[string]int)
	tokenized_texts := make([][]string, len(texts))
	for idx, text := range texts {
		tokenized_texts[idx] = Tokenize(text)
		for _, t := range tokenized_texts[idx] {
			if _, ok := c.WordCounts[t]; !ok {
				c.WordCounts[t] = 0
				c.Vocabulary = append(c.Vocabulary, t)
			}
			c.WordCounts[t] += 1
		}
	}

	parametrized_texts := make([][]int, len(texts))

	for row, tokens := range tokenized_texts {
		parametrized_texts[row] = make([]int, len(c.Vocabulary))
		for _, t := range tokens {
			for col, w := range c.Vocabulary {
				if w == t {
					parametrized_texts[row][col] += 1
				}
			}
		}
	}
	return c.WordCounts, parametrized_texts
}
