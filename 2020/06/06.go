package main

import (
	"bufio"
	"log"
	"os"
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

func countYesQuestionsInGroup(group []string) (int, int) {
	allYesQuestions := 0
	answersMap := map[rune]int{}
	for _, person := range group {
		for _, question := range person {
			answersMap[question] = answersMap[question] + 1
			if answersMap[question] == len(group) {
				allYesQuestions++
			}
		}
	}
	return len(answersMap), allYesQuestions
}

func main() {
	lineBlocks := readLineBlocks("06.txt")

	anyCounts := 0
	everyoneCounts := 0
	for _, lineBlock := range lineBlocks {
		anyCount, everyoneCount := countYesQuestionsInGroup(lineBlock)
		anyCounts += anyCount
		everyoneCounts += everyoneCount
	}

	println("for anyone answered yes", anyCounts)
	println("for everyone answered yes", everyoneCounts)
}
