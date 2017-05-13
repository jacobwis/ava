package ava

import (
	"testing"

	"github.com/wisniewskij26/xu"
)

func TestLookupQValue(t *testing.T) {
	actual := lookupQValue(4, 16)
	expected := 4.05

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}

	actual = lookupQValue(40, 16)
	expected = 5.9
	err = xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}

	actual = lookupQValue(1, 16)
	expected = 3.0
	err = xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}

	actual = lookupQValue(4, 400)
	expected = 3.63
	err = xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}
