package main

import "math"
import "math/rand"
import "slices"
import "os"
import "fmt"
import "strings"
import "strconv"

type Stack struct {
	data []int
	max  int
}

type Equation struct {
    result int
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
	}
}

func Variations(positions int) (result []string) {
	expected_len := int(math.Pow(2, float64(positions)))
	operators := []string{"+", "*"}

	for len(result) != expected_len {
		str := ""
		for range positions {
			str += operators[rand.Intn(2)]
		}
		if !slices.Contains(result, str) {
			result = append(result, str)
		}
	}
	return
}

func IsValid(result int, numbers []int) bool {
	operators_set := Variations(len(numbers) - 1)

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
        for i, m := range members{
            eq.members[i], _ = strconv.Atoi(m) 
        }
        result = append(result, eq)
    }
    return result
}

func main(){
    equations := Load("input")
    part1 := 0

    for i, eq := range equations {
        if IsValid(eq.result, eq.members){
            part1 += eq.result
            fmt.Println(i, "/", len(equations))
        }
    }

    fmt.Println("part1 result is", part1)
}
