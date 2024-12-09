package main

import "testing"
import "reflect"

func TestDistance(t *testing.T) {
	want := 30
	got := Distance("test_input")

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestCoutner(t *testing.T) {
	want := 2
	got := Counter([]int{3, 3}, 3)

	if got != want {
		t.Errorf("wanted %v but got %v", want, got)
	}
}
