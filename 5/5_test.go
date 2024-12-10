package main

import "testing"

func TestIsCorrect(t *testing.T) {
	rules, updates := Load("test_input")

	assertIsCorrect(t, updates[0], rules, 61)
	assertIsCorrect(t, updates[1], rules, 53)
	assertIsCorrect(t, updates[2], rules, 29)
	assertIsCorrect(t, updates[3], rules, 0)
	assertIsCorrect(t, updates[4], rules, 0)
	assertIsCorrect(t, updates[5], rules, 0)
}

func assertIsCorrect(t testing.TB, update string, rules []string, expectation int) {
	got := IsCorrect(update, rules)

	if got != expectation {
		t.Errorf("update %q wants %v but got %v", update, expectation, got)
	}
}
