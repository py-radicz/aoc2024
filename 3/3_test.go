package main

import "testing"

func Test3(t *testing.T) {
	want := 1
	got := 2

	if got != want {
		t.Errorf("wanted %v but got %v", want, got)
	}
}
