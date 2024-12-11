package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	start := []string{"123", "456", "789"}
	want := []string{"123", "456", "789"}

	for range 4 {
		start = Rotate(start)
	}

	if !reflect.DeepEqual(start, want) {
		t.Errorf("wanted %q but got %q instead", want, start)
	}
}

func TestFindGuard(t *testing.T) {
	got := FindGuard([]string{"...", "...", "..^"})
	want := 2

	if got != want {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func TestMoveUp(t *testing.T) {
	got := MoveUp([]string{"...", "...", ".^."})
	want := []string{"...", ".^.", ".X."}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func TestMoveUpWithObstacle(t *testing.T) {
	got := MoveUp([]string{"...", ".#.", ".^."})
	want := Rotate([]string{"...", ".#.", ".^."})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func TestMoveUpFinish(t *testing.T) {
	got := MoveUp([]string{"^..", "...", "..."})
	want := []string{"X..", "...", "..."}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %q but got %q", want, got)
	}
}
