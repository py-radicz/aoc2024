package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func Load(path string) []string {
	data, _ := os.ReadFile(path)
	data_str := strings.TrimSpace(string(data))
	return strings.Split(data_str, "\n")
}

func Transpose(lines []string) []string {
	result := make([]string, len(lines))
	for _, v := range lines {
		for j, char := range v {
			result[j] += string(char)
		}
	}
	return result
}

func Reverse(s string) (result string) {
	for _, r := range s {
		result = string(r) + result
	}
	return
}

func Rotate(matrix []string) (result []string) {
	for i, _ := range matrix {
		var (
			diag  string
			start int = i
			col   int
		)

		for start >= 0 {
			diag += string(matrix[start][col])
			start--
			col++
		}
		result = append(result, diag)
	}

	slices.Reverse(matrix)

	for i, v := range matrix {
		matrix[i] = Reverse(v)
	}

	for i, _ := range matrix {
		var (
			diag  string
			start int = i
			col   int
		)

		for start >= 0 {
			diag += string(matrix[start][col])
			start--
			col++
		}
		result = append(result, Reverse(diag))
	}
	return result[:len(result)-1]
}

func XMASCounter(slice []string) (result int) {
	for _, v := range slice {
		result += strings.Count(v, "XMAS")
		result += strings.Count(v, "SAMX")
	}
	return
}

func SlidingWindows(matrix []string) (result [][]string) {
	for i, row := range matrix {
		if i+2 >= len(matrix) {
			break
		}

		col := 0
		for col < len(row)-3 {
			window := make([]string, 3)
			window[0] = matrix[i][col : col+3]
			window[1] = matrix[i+1][col : col+3]
			window[2] = matrix[i+2][col : col+3]
			result = append(result, window)
			col++
		}
		window := make([]string, 3)
		window[0] = matrix[i][col : col+3]
		window[1] = matrix[i+1][col : col+3]
		window[2] = matrix[i+2][col : col+3]
		result = append(result, window)
	}
	return
}

func GetDiagonal(matrix []string) (result string) {
	for i, row := range matrix {
		for j, _ := range row {
			if i == j {
				result += string(matrix[i][j])
			}
		}
	}
	return
}

func Part2(matrix []string) (result int) {
	for _, window := range SlidingWindows(matrix) {
		var diagonals string = ""
		diagonals += GetDiagonal(window)

		for i, v := range window {
			window[i] = Reverse(v)
		}
		diagonals += GetDiagonal(window)

		if strings.Count(diagonals, "MAS")+strings.Count(diagonals, "SAM") == 2 {
			result += 1
		}
	}
	return
}

func main() {
	lines := Load("input")

	fmt.Printf("%v is part2 result\n", Part2(lines))

	var result int = 0

	result += XMASCounter(lines)
	result += XMASCounter(Transpose(lines))
	result += XMASCounter(Rotate(lines))

	for i, v := range lines {
		lines[i] = Reverse(v)
	}

	result += XMASCounter(Rotate(lines))

	fmt.Printf("%v is part1 result\n", result)
}
