package ava

import "testing"

func TestLookupQValue(t *testing.T) {
	actual := lookupQValue(4, 16)
	expected := 4.05

	if actual != expected {
		t.Errorf("lookupQValue: Expected %v, got %v", expected, actual)
	}

	actual = lookupQValue(40, 16)
	expected = 5.9
	if actual != expected {
		t.Errorf("lookupQValue handles k > 20: Expected %v, got %v", expected, actual)
	}

	actual = lookupQValue(1, 16)
	expected = 3.0
	if actual != expected {
		t.Errorf("lookupQValue handles k < 2: Expected %v, got %v", expected, actual)
	}

	actual = lookupQValue(4, 400)
	expected = 3.63
	if actual != expected {
		t.Errorf("lookupQValue handles df > 120: Expected %v, got %v", expected, actual)
	}
}
