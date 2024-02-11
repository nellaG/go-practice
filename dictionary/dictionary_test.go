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
