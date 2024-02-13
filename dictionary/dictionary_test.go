package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": known}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := known
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, definition)

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := known
	dictionary.Add(word, definition)
	assertDefinition(t, dictionary, word, definition)
}
