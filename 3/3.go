package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Load(path string) []byte {
	data, err := os.ReadFile(path)

	if err != nil {
		return []byte{}
	}

	return data
}

func Matches(data []byte) (result [][]byte) {
	re := regexp.MustCompile(`(?m)mul\(\d{1,3},\d{1,3}\)`)
	result = re.FindAll(data, -1)
	return
}

func main() {
	data := Load("input")

	re := regexp.MustCompile(`(?msU)don\'t\(\).*do\(\)`)
	data = re.ReplaceAll(data, []byte(""))
	matches := Matches(data)
	res := 0

	for _, v := range matches {
		tmp := strings.TrimPrefix(string(v), "mul(")
		tmp = strings.TrimSuffix(tmp, ")")
		vals := strings.Split(tmp, ",")

		part1, _ := strconv.Atoi(vals[0])
		part2, _ := strconv.Atoi(vals[1])

		res += part1 * part2
	}

	fmt.Printf("%v", res)
}
