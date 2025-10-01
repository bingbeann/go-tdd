package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{
		"test": "this is just a test",
	}

	got := Search(dict, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("want %q but got %q given %q", want, got, "test")
	}
}
