package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
)

type Problem struct {
	Question string
	Solution string
}

func main() {

	// Open file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Read CSV file
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problemList := createProblemList(data)
	challenge(problemList)
}

func createProblemList(data [][]string) []Problem {
	var problemList []Problem

	// Store CSV contents into Problem array
	for _, line := range data {
		var problem Problem

		problem.Question = addWhiteSpace(line[0])
		problem.Solution = line[1]

		problemList = append(problemList, problem)
	}

	return problemList
}

func addWhiteSpace(line string) string {
	var whiteSpacedString string
	whiteSpacedString = ""

	// Add whitespace after every character
	for _, character := range line {
		whiteSpacedString = whiteSpacedString + string(character) + " "
	}
	
	return whiteSpacedString
}

func challenge(problemList []Problem) {
	totalQuestions := len(problemList)
	correctQuestions := 0

	for _, problem := range problemList {
		var userInput string

		// Show prompt and get user input
		fmt.Printf("%s:", problem.Question)
		fmt.Scanln(&userInput)

		if userInput != problem.Solution {
			continue
		} else {
			correctQuestions++
		}
	}

	fmt.Printf("Score: %d / %d\n", correctQuestions, totalQuestions)
}

