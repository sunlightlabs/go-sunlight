package sunlight

import (
	"testing"
)

func TestAPIKey(t *testing.T) {
	SetAPIKey("foo")
	foo := GetAPIKey()

	if foo != "foo" {
		t.Fatal("API Key didn't set: " + foo)
	}
}
