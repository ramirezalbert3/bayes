package preprocess

import (
	"testing"
)

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

func check_parametrized_texts_equal(actual [][]float64, expected [][]float64, t *testing.T) {
	if len(actual) != len(expected) {
		t.Error("Expected parametrized_texts len", len(expected), "got", len(actual))
	}
	for i, param_text := range actual {
		if len(param_text) != len(expected[i]) {
			t.Error("Expected text:", i, "len", len(expected[i]), "got", len(param_text))
		}
		for j, count := range param_text {
			if count != expected[i][j] {
				t.Error("Expected count", expected[i][j], "got", count, "at text:", i, "token:", j)
			}
		}
	}
}

func TestCountVectorizerBasic(t *testing.T) {
	texts := []string{"apple banana, bycicle", "bycicle bycicle, blue manycolors and apple"}
	c := CountVectorizer{}
	parametrized_texts := c.FitTransform(texts)

	expected_vocabulary := []string{"apple", "banana", "bycicle", "blue", "manycolors", "and"}
	if len(c.Vocabulary) != len(expected_vocabulary) {
		t.Error("Expected c.Vocabulary len", len(expected_vocabulary), "got", len(c.Vocabulary))
	}
	for idx, word := range c.Vocabulary {
		if expected_vocabulary[idx] != word {
			t.Error("Expected", expected_vocabulary[idx], "got", word)
		}
	}

	expected_word_counts := map[string]int{
		"apple":      2,
		"banana":     1,
		"bycicle":    3,
		"blue":       1,
		"manycolors": 1,
		"and":        1,
	}
	if len(c.WordCounts) != len(expected_word_counts) {
		t.Error("Expected c.WordCounts len", len(expected_word_counts), "got", len(c.WordCounts))
	}
	for word, count := range c.WordCounts {
		if count != expected_word_counts[word] {
			t.Error("Expected count", count, "got", expected_word_counts[word], "for", word)
		}
	}

	expected_parametrized_texts := [][]float64{
		{1, 1, 1, 0, 0, 0},
		{1, 0, 2, 1, 1, 1},
	}
	check_parametrized_texts_equal(parametrized_texts, expected_parametrized_texts, t)

	frequency_texts := TfidfTransform(parametrized_texts, false)

	expected_frequency_texts := [][]float64{
		{1 / 3.0, 1 / 3.0, 1 / 3.0, 0, 0, 0},
		{1 / 6.0, 0, 2 / 6.0, 1 / 6.0, 1 / 6.0, 1 / 6.0},
	}
	check_parametrized_texts_equal(frequency_texts, expected_frequency_texts, t)

	// TODO: test Tfidf Tfidf
}
