package main

import "testing"
import "math"
import "reflect"

func TestVariations(t *testing.T) {
	ops := []string{"+", "*", "|"}
	spaces := 3
	got := Variations(ops, spaces)

	if len(got) != int(math.Pow(float64(len(ops)), float64(spaces))) {
		t.Errorf("bad length of result wanted %v but got %v", math.Pow(float64(len(ops)), float64(spaces)), len(got))
	}
}

func assertStack(t testing.TB, s *Stack, want []int) {
	if !reflect.DeepEqual(s.data, want) {
		t.Errorf("wanted %#v but got %#v instead", want, s.data)
	}
}

func TestStack(t *testing.T) {
	s := Stack{data: []int{}, max: 2}
	assertStack(t, &s, []int{})

	s.Push(1)
	s.Push(2)
	assertStack(t, &s, []int{1, 2})

	res := s.Pop()
	if res != 2 {
		t.Errorf("popped wrong val from stack")
	}
	s.Pop()
	assertStack(t, &s, []int{})

	if s.IsEmpty() != true {
		t.Errorf("wanted true but got false")
	}

	s.Push(1)
	assertStack(t, &s, []int{1})

	if s.IsEmpty() != false {
		t.Errorf("wanted false but got true")
	}

	if s.IsFull() != false {
		t.Errorf("wanted false but got true")
	}

	s.Push(2)
	if s.IsFull() != true {
		t.Errorf("wanted true but got false")
	}

	s.Calc("*")
	assertStack(t, &s, []int{2})
}

func TestEval(t *testing.T) {
	got := Eval([]int{81, 13, 1}, "++")
	want := 95

	if want != got {
		t.Errorf("wanted %v but got %v", want, got)
	}
}
