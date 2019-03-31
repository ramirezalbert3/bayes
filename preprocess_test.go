package preprocess

import (
	"math"
	"reflect"
	"testing"
)

func check_equal(expected, actual interface{}, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Expected:\n", expected, "\nGot:\n", actual)
	}
}

func TestBasic(t *testing.T) {
	test_string := "Hello, my name is Albert, it's 10 times a pleasure to meet you"
	tokens := Tokenize(test_string)
	expected := []string{"hello", "my", "name", "is", "albert", "its", "10", "times", "a", "pleasure", "to", "meet", "you"}
	check_equal(expected, tokens, t)
}

func tfidf(w_in_doc, w_in_docs, doc_len, n_docs int) float64 {
	tf := float64(w_in_doc) / float64(doc_len)
	idf := math.Log(float64(n_docs+1)/float64(w_in_docs+1)) + 1
	return tf * idf
}

func TestCountVectorizerBasic(t *testing.T) {
	texts := []string{"apple banana, bycicle", "bycicle bycicle, blue manycolors and apple"}
	c := CountVectorizer{}
	counts_texts := c.FitTransform(texts)

	expected_vocabulary := []string{"apple", "banana", "bycicle", "blue", "manycolors", "and"}
	check_equal(expected_vocabulary, c.Vocabulary, t)

	expected_word_counts := map[string]int{
		"apple":      2,
		"banana":     1,
		"bycicle":    3,
		"blue":       1,
		"manycolors": 1,
		"and":        1,
	}
	check_equal(expected_word_counts, c.WordCounts, t)

	expected_counts_texts := [][]float64{
		{1, 1, 1, 0, 0, 0},
		{1, 0, 2, 1, 1, 1},
	}
	check_equal(expected_counts_texts, counts_texts, t)

	tf_texts := TfidfTransform(counts_texts, false)

	expected_tf_texts := [][]float64{
		{1 / 3.0, 1 / 3.0, 1 / 3.0, 0, 0, 0},
		{1 / 6.0, 0, 2 / 6.0, 1 / 6.0, 1 / 6.0, 1 / 6.0},
	}
	check_equal(expected_tf_texts, tf_texts, t)

	tfidf_texts := TfidfTransform(counts_texts, true)
	expected_tfidf_texts := [][]float64{
		{tfidf(1, 2, 3, 2), tfidf(1, 1, 3, 2), tfidf(1, 3, 3, 2), 0, 0, 0},
		{tfidf(1, 2, 6, 2), 0, tfidf(2, 3, 6, 2), tfidf(1, 1, 6, 2), tfidf(1, 1, 6, 2), tfidf(1, 1, 6, 2)},
	}
	check_equal(expected_tfidf_texts, tfidf_texts, t)

	// Vars have not been unexpectedly modified
	check_equal([]string{"apple banana, bycicle", "bycicle bycicle, blue manycolors and apple"}, texts, t)
	check_equal(expected_vocabulary, c.Vocabulary, t)
	check_equal(expected_word_counts, c.WordCounts, t)
	check_equal(expected_counts_texts, counts_texts, t)
}
