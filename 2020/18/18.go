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
	Plus       TokenType = 0
	LeftParen  TokenType = 1
	RightParen TokenType = 2
	Multiply   TokenType = 3
	Number     TokenType = 4
)

// Token -
type Token struct {
	tokenType TokenType
	value     string
}

// Stack -
type Stack []Token

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(token Token) {
	*s = append(*s, token)
}

func (s *Stack) pop() Token {
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element
}

func tokenize(expression string) []Token {
	tokenMap := map[TokenType]string{
		Plus:       "^\\+",
		LeftParen:  "^\\(",
		RightParen: "^\\)",
		Multiply:   "^\\*",
		Number:     "^\\d+",
	}

	tokens := make([]Token, 0)
	remainingExpression := strings.ReplaceAll(expression, " ", "")
	for len(remainingExpression) > 0 {
		for key, value := range tokenMap {
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

func evaluateExpression(expression string) int64 {
	tokens := tokenize(expression)
	operatorStack := make([]Token, 0)
	valueStack := make([]int64, 0)

	calcTop := func() {
		operator := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]
		val1, val2 := valueStack[len(valueStack)-2], valueStack[len(valueStack)-1]
		valueStack = valueStack[:len(valueStack)-2]
		if operator.tokenType == Plus {
			valueStack = append(valueStack, val1+val2)
		} else if operator.tokenType == Multiply {
			valueStack = append(valueStack, val1*val2)

		}
	}

	for _, token := range tokens {
		switch token.tokenType {
		case LeftParen:
			println("left paren on stack")

			operatorStack = append(operatorStack, token)
		case Plus:
			// for len(valueStack) > 1 && len(operatorStack) > 0 && (operatorStack[len(operatorStack)-1].tokenType == Plus || operatorStack[len(operatorStack)-1].tokenType == Multiply) {
			// 	number := valueStack[len(valueStack)-1]
			// 	left := valueStack[len(valueStack)-2]

			// 	valueStack = valueStack[:len(valueStack)-2]

			// 	operator := operatorStack[len(operatorStack)-1]

			// 	if operator.tokenType == Multiply {
			// 		operatorStack = operatorStack[:len(operatorStack)-1]

			// 		valueStack = valueStack[:len(valueStack)-1]
			// 		valueStack = append(valueStack, number*left)
			// 		// println("evaluating", number, "*", left)

			// 	} else if operator.tokenType == Plus {
			// 		operatorStack = operatorStack[:len(operatorStack)-1]

			// 		valueStack = valueStack[:len(valueStack)-1]
			// 		valueStack = append(valueStack, number+left)
			// 		// println("evaluating", number, "+", left)

			// 	}
			// }
			operatorStack = append(operatorStack, token)
		case Multiply:
			// for len(valueStack) > 1 && len(operatorStack) > 0 && (operatorStack[len(operatorStack)-1].tokenType == Plus || operatorStack[len(operatorStack)-1].tokenType == Multiply) {
			// 	number := valueStack[len(valueStack)-1]
			// 	valueStack = valueStack[:len(valueStack)-1]

			// 	operator := operatorStack[len(operatorStack)-1]

			// 	if operator.tokenType == Multiply {
			// 		operatorStack = operatorStack[:len(operatorStack)-1]

			// 		left := valueStack[len(valueStack)-1]
			// 		valueStack = valueStack[:len(valueStack)-1]
			// 		valueStack = append(valueStack, number*left)
			// 		// println("evaluating", number, "*", left)

			// 	} else if operator.tokenType == Plus {
			// 		operatorStack = operatorStack[:len(operatorStack)-1]

			// 		left := valueStack[len(valueStack)-1]
			// 		valueStack = valueStack[:len(valueStack)-1]
			// 		valueStack = append(valueStack, number+left)
			// 		// println("evaluating", number, "+", left)

			// 	}
			// }
			operatorStack = append(operatorStack, token)
		case RightParen:
			println("Right paren")
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1].tokenType != LeftParen {
				operator := operatorStack[len(operatorStack)-1]
				// println("OPERATOR FROM STACK", operator.value)
				operatorStack = operatorStack[:len(operatorStack)-1]

				left := valueStack[len(valueStack)-1]
				right := valueStack[len(valueStack)-2]

				valueStack = valueStack[:len(valueStack)-2]

				if operator.tokenType == Multiply {
					// println("evaluating", left, "*", right)
					valueStack = append(valueStack, right*left)
				} else if operator.tokenType == Plus {
					// println("evaluating", left, "+", right)

					valueStack = append(valueStack, right+left)
				}
			}
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1].tokenType == LeftParen {
				// println("LEFT PARAM STOPPINN")
				operatorStack = operatorStack[:len(operatorStack)-1]
			}

			// for len(valueStack) > 1 && len(operatorStack) > 0 && (operatorStack[len(operatorStack)-1].tokenType == Plus || operatorStack[len(operatorStack)-1].tokenType == Multiply) {
			// 	number := valueStack[len(valueStack)-1]
			// 	valueStack = valueStack[:len(valueStack)-1]

			// 	operator := operatorStack[len(operatorStack)-1]

			// 	if operator.tokenType == Multiply {
			// 		operatorStack = operatorStack[:len(operatorStack)-1]

			// 		left := valueStack[len(valueStack)-1]
			// 		valueStack = valueStack[:len(valueStack)-1]
			// 		valueStack = append(valueStack, number*left)
			// 		// println("evaluating", number, "*", left)

			// 	} else if operator.tokenType == Plus {
			// 		operatorStack = operatorStack[:len(operatorStack)-1]

			// 		left := valueStack[len(valueStack)-1]
			// 		valueStack = valueStack[:len(valueStack)-1]
			// 		valueStack = append(valueStack, number+left)
			// 		// println("evaluating", number, "+", left)

			// 	}
			// }
		case Number:
			numberS, _ := strconv.Atoi(token.value)
			number := int64(numberS)
			if len(operatorStack) == 0 {
				valueStack = append(valueStack, number)
			} else {
				operator := operatorStack[len(operatorStack)-1]

				if operator.tokenType == Multiply {
					operatorStack = operatorStack[:len(operatorStack)-1]

					left := valueStack[len(valueStack)-1]
					valueStack = valueStack[:len(valueStack)-1]
					valueStack = append(valueStack, number*left)
					// println("evaluating", number, "*", left)

				} else if operator.tokenType == Plus {
					operatorStack = operatorStack[:len(operatorStack)-1]

					left := valueStack[len(valueStack)-1]
					valueStack = valueStack[:len(valueStack)-1]
					valueStack = append(valueStack, number+left)
					// println("evaluating", number, "+", left)

				} else {
					valueStack = append(valueStack, number)
				}
			}
		}
	}

	for len(operatorStack) != 0 {
		calcTop()
	}

	return valueStack[0]
}

func part1(lines []string) int64 {
	var sum int64 = 0
	for _, line := range lines {
		sum += evaluateExpression(line)
	}
	return sum
}

func part2(lines []string) int {
	return 0
}

func main() {
	lines := readLines("18.txt")
	println(part1(lines))
	println(part2(lines))
}
