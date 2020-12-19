package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLineBlocks(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineBlocks := make([][]string, 0)
	currentLineBlock := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			lineBlocks = append(lineBlocks, currentLineBlock)
			currentLineBlock = make([]string, 0)
		} else {
			currentLineBlock = append(currentLineBlock, line)
		}
	}
	lineBlocks = append(lineBlocks, currentLineBlock)

	return lineBlocks
}

// Rule -
type Rule struct {
	subrules [][]int
	pattern  string
}

func parseRules(lines []string) []*Rule {
	rules := make([]*Rule, len(lines))
	for _, line := range lines {
		split := strings.Split(line, ": ")
		ruleIndex, _ := strconv.Atoi(split[0])

		subrulesSplit := strings.Split(split[1], " | ")
		if strings.HasPrefix(subrulesSplit[0], "\"") {
			rules[ruleIndex] = &Rule{pattern: strings.ReplaceAll(subrulesSplit[0], "\"", "")}
		} else {
			subrules := make([][]int, 0)

			for _, subruleText := range subrulesSplit {
				subrule := make([]int, 0)
				subruleSplit := strings.Split(subruleText, " ")

				for _, subruleIndex := range subruleSplit {
					idx, _ := strconv.Atoi(subruleIndex)
					subrule = append(subrule, idx)
				}
				subrules = append(subrules, subrule)
			}
			rules[ruleIndex] = &Rule{subrules: subrules}

		}

	}
	return rules
}

func buildRegex(rules []*Rule, ruleIdx int, recursionDepth int, recursionLimit int) string {
	if rules[ruleIdx].pattern != "" {
		return rules[ruleIdx].pattern
	}

	result := "(?:"

	for i, subrule := range rules[ruleIdx].subrules {
		for _, rule := range subrule {
			if recursionDepth < recursionLimit {
				result += buildRegex(rules, rule, recursionDepth+1, recursionLimit)
			} else {
				result += ".+"
			}
		}
		if i < len(rules[ruleIdx].subrules)-1 {
			result += "|"
		}
	}

	return result + ")"
}

func matchMessages(rules []*Rule, messages []string) int {
	matched := 0
	regex := regexp.MustCompile("^" + buildRegex(rules, 0, 0, len(rules)) + "$")
	for _, message := range messages {
		if regex.MatchString(message) {
			matched++
		}
	}
	return matched
}

func part1(rules []*Rule, messages []string) int {
	return matchMessages(rules, messages)
}

func part2(rules []*Rule, messages []string) int {
	rules[8].subrules = append(rules[8].subrules, []int{42, 8})
	rules[11].subrules = append(rules[11].subrules, []int{42, 11, 31})

	return matchMessages(rules, messages)
}

func main() {
	blocks := readLineBlocks("19.txt")
	rules := parseRules(blocks[0])
	println(part1(rules, blocks[1]))
	println(part2(rules, blocks[1]))
}
