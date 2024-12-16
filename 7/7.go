package main

import "math"
import "math/rand"
import "slices"
import "os"
import "fmt"
import "strings"
import "strconv"

var cache = map[int][]string{}

type Stack struct {
	data []int
	max  int
}

type Equation struct {
	result  int
	members []int
}

func (s *Stack) Push(num int) {
	s.data = append(s.data, num)
}

func (s *Stack) Pop() int {
	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return result
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) IsFull() bool {
	return len(s.data) == s.max
}

func (s *Stack) Calc(op string) {
	switch op {
	case "+":
		s.Push(s.Pop() + s.Pop())
	case "*":
		s.Push(s.Pop() * s.Pop())
	case "|":
		second := strconv.Itoa(s.Pop())
		first := strconv.Itoa(s.Pop())
		val, _ := strconv.Atoi(first + second)
		s.Push(val)
	}
}

// bruteforce
func Variations(operators []string, positions int) (result []string) {
	expected_len := int(math.Pow(float64(len(operators)), float64(positions)))

	if val, ok := cache[positions]; ok {
		return val
	}

	if positions == 11 {
		for _, str := range cache[10] {
			for _, op := range operators {
				result = append(result, str+op)
			}
		}
	}

	for len(result) != expected_len {
		str := ""
		for range positions {
			str += operators[rand.Intn(len(operators))]
		}
		if !slices.Contains(result, str) {
			result = append(result, str)
		}
	}
	cache[positions] = result
	return
}

func IsValid(result int, numbers []int, operators []string) bool {
	operators_set := Variations(operators, len(numbers)-1)

	for _, operators := range operators_set {
		if result == Eval(numbers, operators) {
			return true
		}
	}
	return false
}

func Eval(nums []int, operators string) int {
	s := Stack{data: []int{}, max: 2}
	op := 0

	for _, num := range nums {
		if s.IsFull() {
			s.Calc(string(operators[op]))
			op++
			s.Push(num)
		} else {
			s.Push(num)
		}
	}
	s.Calc(string(operators[op]))
	return s.Pop()
}

func Load(path string) (result []Equation) {
	data, _ := os.ReadFile(path)
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, row := range rows {
		eq := Equation{}
		divided := strings.Split(row, ":")
		members := strings.Fields(divided[1])
		eq.result, _ = strconv.Atoi(divided[0])
		eq.members = make([]int, len(members))
		for i, m := range members {
			eq.members[i], _ = strconv.Atoi(m)
		}
		result = append(result, eq)
	}
	return result
}

func MathResolver(equations []Equation, operators []string) (result int) {
	for _, eq := range equations {
		if IsValid(eq.result, eq.members, operators) {
			result += eq.result
		}
	}
	return
}

func main() {
	equations := Load("test_input")
	fmt.Println("part1 result is", MathResolver(equations, []string{"+", "*"}))
	cache = map[int][]string{}
	fmt.Println("part2 result is", MathResolver(equations, []string{"+", "*", "|"}))
}
