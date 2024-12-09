package main

import "testing"

func TestIsSafe(t *testing.T) {
	want := true
	got := IsSafe([]int{22, 25, 27, 29, 30, 31, 33})

	if want != got {
		t.Errorf("wanted %v but got %v instead", want, got)
	}
}

func TestIsSafeDampened(t *testing.T) {
	want := true
	got := IsSafeDampened([]int{4, 7, 8, 9})

	if want != got {
		t.Errorf("wanted %v but got %v instead", want, got)
	}
}
