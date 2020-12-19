package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// TokenType -
type TokenType int

// -
const (
	Plus TokenType = iota
	LeftParen
	RightParen
	Multiply
	Number
)

// Token -
type Token struct {
	tokenType TokenType
	value     string
}

var tokenRegexMap = map[TokenType]string{
	Plus:       "^\\+",
	LeftParen:  "^\\(",
	RightParen: "^\\)",
	Multiply:   "^\\*",
	Number:     "^\\d+",
}

func tokenize(expression string) []Token {
	tokens := make([]Token, 0)
	remainingExpression := strings.ReplaceAll(expression, " ", "")

	for len(remainingExpression) > 0 {
		for key, value := range tokenRegexMap {
			regex, _ := regexp.Compile(value)
			loc := regex.FindStringIndex(remainingExpression)

			if len(loc) > 0 && loc[1] > 0 {
				token := Token{tokenType: key, value: remainingExpression[loc[0]:loc[1]]}
				tokens = append(tokens, token)
				remainingExpression = remainingExpression[loc[1]:]
			}
		}
	}
	return tokens
}

func applyOperator(operator TokenType, left int, right int) int {
	if operator == Plus {
		return left + right
	} else if operator == Multiply {
		return left * right
	}
	return 0
}

func evaluateExpression(expression string, precendence map[TokenType]int) int {
	tokens := tokenize(expression)

	// why don't you have generics Go...
	operatorStack := make([]Token, 0)
	valueStack := make([]int, 0)

	applyOperators := func() {
		operator := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]

		left, right := valueStack[len(valueStack)-2], valueStack[len(valueStack)-1]
		valueStack = valueStack[:len(valueStack)-2]

		valueStack = append(valueStack, applyOperator(operator.tokenType, left, right))
	}

	canApplyOperator := func(tokenType TokenType) bool {
		if len(valueStack) < 1 || len(operatorStack) == 0 {
			return false
		}

		lastOperator := operatorStack[len(operatorStack)-1].tokenType
		if lastOperator != Plus && lastOperator != Multiply {
			return false
		}

		return precendence[operatorStack[len(operatorStack)-1].tokenType] >= precendence[tokenType]
	}

	for _, token := range tokens {
		switch token.tokenType {
		case LeftParen:
			operatorStack = append(operatorStack, token)
		case Plus:
			for canApplyOperator(Plus) {
				applyOperators()
			}
			operatorStack = append(operatorStack, token)
		case Multiply:
			for canApplyOperator(Multiply) {
				applyOperators()
			}
			operatorStack = append(operatorStack, token)
		case RightParen:
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1].tokenType != LeftParen {
				applyOperators()
			}
			// Pop left parenthesis
			operatorStack = operatorStack[:len(operatorStack)-1]
		case Number:
			number, _ := strconv.Atoi(token.value)
			valueStack = append(valueStack, number)
		}
	}

	for len(operatorStack) > 0 {
		applyOperators()
	}

	return valueStack[0]
}

func sumResults(lines []string, precendence map[TokenType]int) int64 {
	var sum int64 = 0
	for _, line := range lines {
		sum += int64(evaluateExpression(line, precendence))
	}
	return sum
}

func part1(lines []string) int64 {
	return sumResults(lines, map[TokenType]int{Plus: 1, Multiply: 1})
}

func part2(lines []string) int64 {
	return sumResults(lines, map[TokenType]int{Plus: 2, Multiply: 1})
}

func main() {
	lines := readLines("18.txt")
	println(part1(lines))
	println(part2(lines))
}
