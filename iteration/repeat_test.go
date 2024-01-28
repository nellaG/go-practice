package iteration

import "testing"

func TestRepeat(t *testing.T) {
	// t.Fatal("not implemented")
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but %q", expected, repeated)
	}
}
