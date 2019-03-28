package bayes

import "strings" // TODO: regex

func tokenize(s string) []string {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "'", "", -1)
	s = strings.ToLower(s)
	tokens := strings.Split(s, " ")
	return tokens
}
