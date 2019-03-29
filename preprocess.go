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
	tokens := make([][]string, len(texts))
	for idx, text := range texts {
		tokens[idx] = Tokenize(text)
		for _, t := range tokens[idx] {
			if _, ok := c.WordCounts[t]; !ok {
				c.WordCounts[t] = 0
				c.Vocabulary = append(c.Vocabulary, t)
			}
			c.WordCounts[t] += 1
		}
	}

	counts := make([][]int, len(texts))

	for t_idx, text_tokens := range tokens {
		counts[t_idx] = make([]int, len(c.Vocabulary))
		for _, t := range text_tokens {
			for w_idx, w := range c.Vocabulary {
				if w == t {
					counts[t_idx][w_idx] += 1
				}
			}
		}
	}
	return c.WordCounts, counts
}
