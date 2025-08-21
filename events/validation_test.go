package events

import "testing"

func TestIsValidTitle(t *testing.T) {
	result := isValidTitle("Hello World")
	if !result {
		t.Error(
			"For 'Hello World' expected true, got false",
		)
	}

	result = isValidTitle("")
	if result {
		t.Error(
			"For '' expected false, got true",
		)
	}

}
