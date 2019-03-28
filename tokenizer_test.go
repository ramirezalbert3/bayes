package bayes

import "testing"

func TestBasic(t *testing.T) {
	test_string := "Hello, my name is Albert, it's a pleasure to meet you"
	tokens := tokenize(test_string)
	expected := []string{"hello", "my", "name", "is", "albert", "its", "a", "pleasure", "to", "meet", "you"}
	for idx, token := range tokens {
		if token != expected[idx] {
			t.Error("Expected", expected[idx], "got", token)
		}
	}
}
