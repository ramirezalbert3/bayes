package tokenizer

import "testing"

func TestBasic(t *testing.T) {
	test_string := "Hello, my name is Albert, it's 10 times a pleasure to meet you"
	tokens := Tokenize(test_string)
	expected := []string{"hello", "my", "name", "is", "albert", "its", "10", "times", "a", "pleasure", "to", "meet", "you"}
	for idx, token := range tokens {
		if token != expected[idx] {
			t.Error("Expected", expected[idx], "got", token)
		}
	}
}

func TestKeepsUnderscoresAndDashes(t *testing.T) {
	test_string := "This... sho.uld-be fai//Rly_siMple"
	tokens := Tokenize(test_string)
	expected := []string{"this", "should-be", "fairly_simple"}
	for idx, token := range tokens {
		if token != expected[idx] {
			t.Error("Expected", expected[idx], "got", token)
		}
	}
}
