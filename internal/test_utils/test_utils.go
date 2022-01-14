package test_utils

import "testing"

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func AssertError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("missing error: %s", err)
	}
}
