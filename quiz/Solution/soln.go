package main

import(
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Solution string
}

func main() {
	// Declare the flag/switches
	csvFileName := flag.String("csv", "problems.csv", "A CSV file in the format  of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")

	// Parse command line into defined flags
	flag.Parse()

	// Create file handle
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open CSV file %s\n", *csvFileName)
		os.Exit(1)
	}

	defer file.Close()
	
	// Open CSV reader
	r := csv.NewReader(file)

	// Read in all lines
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read CSV file %s\n", *csvFileName)
		os.Exit(1)
	}

	// Convert lines to slice
	problems := parseLines(lines)
	
	// Init timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	
	correct := 0

	problemLoop:
	for i, p := range problems {
		// Stdout problem
		fmt.Printf("Problem #%d: %s = ", i+1, p.Question)

		answerCh := make(chan string)
		go func() {
			// Stdin user input
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer // Send user input to answer channel
		}() // Put `()` at the end to call this immediatly

		select {
		case <-timer.C: // If you get message from timer channel
			fmt.Println()
			break problemLoop
		case answer := <- answerCh: // If you get an answer from the answer channel
			if answer == p.Solution {
				correct++
			}
		}
	}

	fmt.Printf("\nYou scored %d / %d\n", correct, len(problems))
}

func parseLines(lines [][]string) []Problem {
	// Declare return value
	ret := make([]Problem, len(lines))

	// Create slice of Problem structs and fill in Question and Solutions
	// TIP: When you know the size of something, there shouldn't be a reason to use append.
	//      Allocate the needed memory and fill in the values.
	for i, line := range lines {
		ret[i] = Problem{
			Question: line[0],
			Solution: strings.TrimSpace(line[1]),
		}
	}

	return ret
}