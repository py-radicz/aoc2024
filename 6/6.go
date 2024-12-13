package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type GuardPos struct {
	X int
	Y int
}

var rotate_counter int = 0

func Load(path string) [][]string {
	data, _ := os.ReadFile(path)
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	matrix := make([][]string, len(rows))

	for i, row := range rows {
		matrix[i] = make([]string, len(row))
		for j, col := range row {
			matrix[i][j] = string(col)
		}
	}
	return matrix
}

func Rotate(matrix [][]string) {
	tmp := make([][]string, len(matrix))

	for i, row := range matrix {
		tmp[i] = make([]string, len(row))
		for j, _ := range row {
			tmp[i][j] = matrix[j][i]
		}
	}

	for i, row := range tmp {
		for j, _ := range row {
			matrix[i][j] = tmp[i][j]
		}
	}
	slices.Reverse(matrix)
	rotate_counter++
}

func FindGuard(matrix [][]string) GuardPos {
	for i, row := range matrix {
		for j, _ := range row {
			if matrix[i][j] == "^" {
				return GuardPos{X: i, Y: j}
			}
		}
	}
	return GuardPos{X: -1, Y: -1}
}

func MoveUp(matrix [][]string) {
	guard := FindGuard(matrix)

	if guard.X == 0 {
		matrix[guard.X][guard.Y] = "X"
		return
	}

	if matrix[guard.X-1][guard.Y] == "#" {
		Rotate(matrix)
		return
	}

	matrix[guard.X][guard.Y] = "X"
	matrix[guard.X-1][guard.Y] = "^"

	return
}

func VisitedPosCounter(matrix [][]string) (result int) {
	for _, row := range matrix {
		for _, col := range row {
			if col == "X" {
				result += 1
			}
		}
	}
	return
}

func GuardRouteSteps(matrix [][]string) int {
	steps := 0
	for FindGuard(matrix).X >= 0 {
		steps++
		MoveUp(matrix)

		if steps > 10000 {
			return -1
		}
	}

	return VisitedPosCounter(matrix)
}

func GetVisited(matrix [][]string) []GuardPos {
	// reset to original orientation
	for rotate_counter%4 != 0 {
		Rotate(matrix)
	}

	positions := []GuardPos{}

	for i, row := range matrix {
		for j, _ := range row {
			if matrix[i][j] == "X" {
				positions = append(positions, GuardPos{X: i, Y: j})
			}
		}
	}
	return positions

}

func DisplayMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
}

func main() {
	filename := "input"
	matrix := Load(filename)
	start_pos := FindGuard(matrix)
	result1 := GuardRouteSteps(matrix)
	path := GetVisited(matrix)

	fmt.Printf("part1 result: %v, visited positions: %v, guard start_pos: %v\n", result1, len(path), start_pos)
	DisplayMatrix(matrix)

	var loops int
	for _, pos := range path {
		if pos == start_pos {
			continue
		}
		matrix = Load(filename)
		matrix[pos.X][pos.Y] = "#"

		if GuardRouteSteps(matrix) == -1 {
			loops++
			fmt.Printf("%v loop detected with obstacle on\n", pos)
		}

	}

	fmt.Printf("part2 result: %v", loops)
	// took hint from reddit that obstacle may only appear on Guard usual route
}
