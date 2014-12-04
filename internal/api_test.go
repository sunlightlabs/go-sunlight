package internal

import "testing"

func TestGenerateURL(t *testing.T) {
	url := GenerateURL("https://fnord", "foo", "bar")
	if url != "https://fnord/foo/bar" {
		t.Fatal("URL invalid: " + url)
	}
}
