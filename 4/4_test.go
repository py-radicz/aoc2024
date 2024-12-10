package main

import "testing"
import "reflect"

func TestTranspose(t *testing.T) {
	got := Transpose([]string{"abc", "def", "ghi"})
	want := []string{"adg", "beh", "cfi"}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func TestRotate(t *testing.T) {
	got := Rotate([]string{"abc", "def", "ghi"})
	want := []string{"a", "db", "gec", "i", "hf"}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func TestPart1(t *testing.T) {
	matrix := Load("test_input")
	result := 0

	result += XMASCounter(matrix)
	result += XMASCounter(Transpose(matrix))
	result += XMASCounter(Rotate(matrix))

	for i, v := range matrix {
		matrix[i] = Reverse(v)
	}

	result += XMASCounter(Rotate(matrix))

	want := 18

	if want != result {
		t.Errorf("wanted %v but got %v", want, result)
	}
}

func TestReverse(t *testing.T) {
	got := Reverse("Radim")
	want := "midaR"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}

}

func TestSlidingWindows(t *testing.T) {
	got := SlidingWindows([]string{"abcd", "efgh", "ijkl"})
	want := [][]string{
		{"abc", "efg", "ijk"},
		{"bcd", "fgh", "jkl"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %q but got %q instead", want, got)
	}
}

func TestGetDiagonal(t *testing.T) {
	got := GetDiagonal([]string{"abc", "def", "ghi"})
	want := "aei"

	if got != want {
		t.Errorf("wanted %q but got %q instead", want, got)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(Load("test_input"))
	want := 9

	if got != want {
		t.Errorf("wanted %v but got %v", want, got)
	}
}
