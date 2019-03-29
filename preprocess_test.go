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
	c := CountVectorizer{}
	word_counts, parametrized_texts := c.FitTransform(texts)
	expected_word_counts := map[string]int{
		"apple":      2,
		"banana":     1,
		"bycicle":    3,
		"blue":       1,
		"manycolors": 1,
		"and":        1,
	}
	if len(word_counts) != len(expected_word_counts) {
		t.Error("Expected word_counts len", len(expected_word_counts), "got", len(word_counts))
	}
	for word, count := range word_counts {
		if count != expected_word_counts[word] {
			t.Error("Expected count", count, "got", expected_word_counts[word], "for", word)
		}
	}

	expected_parametrized_texts := [][]int{
		{1, 1, 1, 0, 0, 0},
		{1, 0, 2, 1, 1, 1},
	}
	if len(parametrized_texts) != len(expected_parametrized_texts) {
		t.Error("Expected parametrized_texts len", len(expected_parametrized_texts), "got", len(parametrized_texts))
	}
	for i, param_text := range parametrized_texts {
		if len(param_text) != len(expected_parametrized_texts[i]) {
			t.Error("Expected text:", i, "len", len(expected_parametrized_texts[i]), "got", len(param_text))
		}
		for j, count := range param_text {
			if count != expected_parametrized_texts[i][j] {
				t.Error("Expected count", expected_parametrized_texts[i][j], "got", count, "for", i, j)
			}
		}
	}
}
