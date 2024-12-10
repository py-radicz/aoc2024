package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Load(path string) (rules, updates []string) {
	data, _ := os.ReadFile(path)
	instructions := strings.Split(string(data), "\n\n")

	for _, rule := range strings.Split(instructions[0], "\n") {
		rules = append(rules, rule)
	}

	for _, update := range strings.Split(strings.TrimSpace(instructions[1]), "\n") {
		updates = append(updates, update)
	}

	return
}

func IsCorrect(update string, rules []string) int {
	updates := strings.Split(update, ",")

	for _, up := range updates {
		for _, rule := range rules {
			if strings.HasPrefix(rule, string(up)+"|") {
				up_pos := strings.Index(update, up)
				rule_pos := strings.Index(update, strings.Split(rule, "|")[1])

				if rule_pos == -1 {
					continue
				}

				if up_pos > rule_pos {
					return 0
				}
			}
		}
	}
	i, _ := strconv.Atoi(updates[len(updates)/2])
	return i

}

func CorrectUpdate(update string, rules []string) (result int) {
	updates := strings.Split(update, ",")

	for IsCorrect(strings.Join(updates, ","), rules) == 0 {
		for _, up := range updates {
			for _, rule := range rules {
				if strings.HasPrefix(rule, string(up)+"|") {
					up_pos := slices.Index(updates, up)
					rule_pos := slices.Index(updates, strings.Split(rule, "|")[1])

					if rule_pos == -1 {
						continue
					}
					if up_pos > rule_pos {
						updates[up_pos], updates[rule_pos] = updates[rule_pos], updates[up_pos]
					}
				}
			}
		}
	}
	return IsCorrect(strings.Join(updates, ","), rules)
}

func main() {
	rules, updates := Load("input")
	var (
		result  int
		result2 int
	)

	for _, up := range updates {
		result += IsCorrect(up, rules)
	}
	fmt.Printf("part1 result is %v\n", result)

	for _, up := range updates {
		if IsCorrect(up, rules) == 0 {
			result2 += CorrectUpdate(up, rules)
		}
	}
	fmt.Printf("part2 result is %v\n", result2)
}
