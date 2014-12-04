package internal

import "testing"

func TestGenerateURL(t *testing.T) {
	url := GenerateURL("https://fnord", "foo", "bar")
	if url != "https://fnord/foo/bar" {
		t.Fatal("URL invalid: " + url)
	}
}

func TestQueryURI(t *testing.T) {
	url := QueryURI("https://fnord", map[string]string{
		"baz": "111",
	}, "foo")

	if url != "https://fnord/foo?apikey=foo&baz=111" {
		t.Fatal("URL invalid: " + url)
	}
}
