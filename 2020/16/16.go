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
	name   string
	ranges [2][2]int
}

func parseRange(start string, end string) [2]int {
	startNumber, _ := strconv.Atoi(start)
	endNumber, _ := strconv.Atoi(end)
	return [2]int{startNumber, endNumber}
}

func parseTicket(text string) []int {
	split := strings.Split(text, ",")
	numbers := make([]int, len(split))
	for i := 0; i < len(split); i++ {
		number, _ := strconv.Atoi(split[i])
		numbers[i] = number
	}
	return numbers
}

func parseInput(lines [][]string) ([]Rule, []int, [][]int) {
	rules := make([]Rule, 0)
	for _, ruleText := range lines[0] {
		regexp := regexp.MustCompile(`([a-z\ ]+): (\d+)-(\d+) or (\d+)-(\d+)`)
		matches := regexp.FindStringSubmatch(ruleText)

		rules = append(rules, Rule{name: matches[1], ranges: [2][2]int{
			parseRange(matches[2], matches[3]),
			parseRange(matches[4], matches[5]),
		}})
	}

	myTicket := parseTicket(lines[1][1])

	otherTickets := make([][]int, 0)
	for i := 1; i < len(lines[2]); i++ {
		otherTickets = append(otherTickets, parseTicket(lines[2][i]))
	}

	return rules, myTicket, otherTickets
}

func validateRule(ticketField int, rule Rule) bool {
	for _, ruleRange := range rule.ranges {
		if ticketField >= ruleRange[0] && ticketField <= ruleRange[1] {
			return true
		}
	}
	return false
}

func part1(rules []Rule, tickets [][]int) (int, [][]int) {
	errorRate := 0
	validTickets := make([][]int, 0)

	for _, ticket := range tickets {
		ticketValid := true
		for _, ticketField := range ticket {
			anyRuleValid := false
			for _, rule := range rules {
				if validateRule(ticketField, rule) {
					anyRuleValid = true
					continue
				}
			}
			if !anyRuleValid {
				ticketValid = false
				errorRate += ticketField
			}
		}
		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	return errorRate, validTickets
}

func part2(rules []Rule, tickets [][]int, myTicket []int) int64 {
	println(len(tickets))
	validRules := make([][]Rule, len(tickets[0]))

	for i := 0; i < len(tickets[0]); i++ {
		validRules[i] = make([]Rule, 0)
		for _, rule := range rules {
			valid := true
			for _, ticket := range tickets {
				if !validateRule(ticket[i], rule) {
					valid = false
					continue
				}
			}
			if valid {
				validRules[i] = append(validRules[i], rule)
			}
		}
	}

	// this should really be solved with a perfect matching algorithm, but for
	// this task the greedy approach seems to work
	var result int64 = 1
	usedRules := map[string]bool{}
	for i := 0; i < len(tickets[0]); i++ {
		for j := 0; j < len(tickets[0]); j++ {
			stillValid := 0
			var firstValid Rule

			for k := 0; k < len(validRules[j]); k++ {
				_, ok := usedRules[validRules[j][k].name]
				if !ok {
					firstValid = validRules[j][k]
					stillValid++
				}
			}
			if stillValid == 1 {
				usedRules[firstValid.name] = true
				if strings.HasPrefix(firstValid.name, "departure") {
					result *= int64(myTicket[j])
				}
				continue
			}
		}
	}

	return result
}

func main() {
	lines := readLineBlocks("16.txt")
	rules, myTicket, nearbyTickets := parseInput(lines)
	errorRate, validTickets := part1(rules, nearbyTickets)
	println(errorRate)
	println(part2(rules, validTickets, myTicket))
}
