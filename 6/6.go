package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func Load(path string) []string {
	data, _ := os.ReadFile(path)
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func Rotate(lines []string) []string {
	result := make([]string, len(lines))
	for _, v := range lines {
		for j, char := range v {
			result[j] += string(char)
		}
	}

	slices.Reverse(result)
	return result
}

func FindGuard(matrix []string) int {
	for i, row := range matrix {
		if strings.Contains(row, "^") {
			return i
		}
	}
	return -1
}

func MoveUp(matrix []string) []string {
	guard_row := FindGuard(matrix)
	guard_col := strings.Index(matrix[guard_row], "^")

	if guard_row == 0 {
		str := []rune(matrix[guard_row])
		str[guard_col] = 'X'
		matrix[guard_row] = string(str)
		return matrix
	}

	if matrix[guard_row-1][guard_col] == '#' {
		return Rotate(matrix)
	}

	str := []rune(matrix[guard_row-1])
	str[guard_col] = '^'
	matrix[guard_row-1] = string(str)
	matrix[guard_row] = strings.Replace(matrix[guard_row], "^", "X", 1)
	return matrix
}

func VisitedPosCounter(matrix []string) (result int) {
	for _, row := range matrix {
		result += strings.Count(row, "X")
	}
	return
}

func main() {
	matrix := Load("input")

	for FindGuard(matrix) >= 0 {
		matrix = MoveUp(matrix)
	}

	fmt.Printf("part 1 result is: %v", VisitedPosCounter(matrix))
}
