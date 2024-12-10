package main

import (
	"fmt"
	"os"
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

func main() {
	rules, updates := Load("input")
    var result int

    for _, up := range updates {
        result += IsCorrect(up, rules)
    }
	fmt.Printf("part1 result is %v\n", result)
}
