package main

import "testing"

var value = "this is just a test"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": value}

	got := Search(dict, "test")
	want := value

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
