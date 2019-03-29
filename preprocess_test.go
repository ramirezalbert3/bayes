package preprocess

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


func TestCountVectorizerBasic(t *testing.T) {
	texts := []string{"apple banana, bycicle", "bycicle bycicle, blue manycolors and apple"}
	vocabulary := CountVectorizer(texts)
	expected := map[string]int{
		"apple": 2,
		"banana": 1,
		"bycicle": 3,
		"blue": 1,
		"manycolors": 1,
		"and": 1,
	}
	for word, count := range vocabulary {
		if count != expected[word] {
			t.Error("Expected count", count, "got", expected[word], "for", word)
		}
	}
}