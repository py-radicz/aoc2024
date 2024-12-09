package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Load(path string) (result [][]int) {
	file, err := os.Open(path)

	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		tmp := []int{}

		for _, v := range line {
			num, _ := strconv.Atoi(v)
			tmp = append(tmp, num)
		}
		result = append(result, tmp)
	}
	return

}

func IsSafe(slice []int) (result bool) {
	tmp := slice[0]
	for i, v := range slice {
		if i == 0 {
			continue
		}

		if v > tmp && v-tmp >= 1 && v-tmp <= 3 {
			result = true
			tmp = v
		} else {
			result = false
			break
		}
	}

	if result {
		return result
	}

	tmp = slice[0]
	for i, v := range slice {
		if i == 0 {
			continue
		}

		if v < tmp && tmp-v >= 1 && tmp-v <= 3 {
			result = true
			tmp = v
		} else {
			result = false
			break
		}
	}

	return result
}

func IsSafeDampened(slice []int) bool {
	for i, _ := range slice {
		list := []int{}

		for j, _ := range slice {
			if j != i {
				list = append(list, slice[j])
			}
		}
		if IsSafe(list) {
			return true
		}
	}
	return false
}

func main() {
	data := Load("input")
	sum := 0

	for _, sl := range data {
		if IsSafe(sl) {
			sum += 1
		} else {
			if IsSafeDampened(sl) {
				sum += 1
			}
		}
	}
	fmt.Printf("total safe sequences: %v", sum)
}

// seems it should be >531 but <539
