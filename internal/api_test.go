package internal

import (
	"github.com/sunlightlabs/go-sunlight/sunlight"
	"testing"
)

func TestGenerateURL(t *testing.T) {
	url := GenerateURL("https://fnord", "foo", "bar")
	if url != "https://fnord/foo/bar" {
		t.Fatal("URL invalid: " + url)
	}
}

func TestQueryURL(t *testing.T) {
	sunlight.SetAPIKey("example")

	url := QueryURL("https://fnord", map[string]string{
		"baz": "111",
	}, "foo")

	if url != "https://fnord/foo?apikey=example&baz=111" {
		t.Fatal("URL invalid: " + url)
	}
}
