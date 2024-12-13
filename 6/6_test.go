package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	start := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	want := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	for range 4 {
		Rotate(start)
	}

	if !reflect.DeepEqual(start, want) {
		t.Errorf("wanted %q but got %q instead", want, start)
	}
}

func TestFindGuard(t *testing.T) {
	matrix := [][]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "^"},
	}
	got := FindGuard(matrix)
	want := GuardPos{X: 2, Y: 2}

	if got != want {
		t.Errorf("wanted %v but got %v", want, got)
	}

	matrix[2][2] = "."
	got = FindGuard(matrix)
	want = GuardPos{X: -1, Y: -1}

	if got != want {
		t.Errorf("wanted %v but got %v", want, got)
	}

}

func TestMoveUp(t *testing.T) {
	matrix := [][]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", "^", "."},
	}
	MoveUp(matrix)
	want := [][]string{
		{".", ".", "."},
		{".", "^", "."},
		{".", "X", "."},
	}

	if !reflect.DeepEqual(matrix, want) {
		t.Errorf("wanted %q but got %q", want, matrix)
	}
}

func TestMoveUpWithObstacle(t *testing.T) {
	matrix := [][]string{
		{".", ".", "."},
		{".", "#", "."},
		{".", "^", "."},
	}
	MoveUp(matrix)
	want := [][]string{
		{".", ".", "."},
		{".", "#", "^"},
		{".", ".", "."},
	}

	if !reflect.DeepEqual(matrix, want) {
		t.Errorf("wanted %q but got %q", want, matrix)
	}
}

func TestMoveUpFinish(t *testing.T) {
	matrix := [][]string{
		{"^", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}
	MoveUp(matrix)
	want := [][]string{
		{"X", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}

	if !reflect.DeepEqual(matrix, want) {
		t.Errorf("wanted %q but got %q", want, matrix)
	}
}
