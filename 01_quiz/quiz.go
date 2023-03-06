package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	csv   string
	limit int
}

type Problem struct {
	question string
	answer   string
}

func parseConfig() Config {
	csv := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	return Config{*csv, *limit}
}

func handleError(e error, msg string) {
	if e != nil {
		fmt.Printf(msg)
		os.Exit(1)
	}
}

func readFile(filename string) [][]string {
	file, err := os.Open(filename)
	handleError(err, fmt.Sprintf("Failed to open the CSV file: %s\n", filename))

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	handleError(err, fmt.Sprintf("Failed to read the CSV file: %s\n", filename))

	return lines
}

func parseQuestions(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))

	for i, line := range lines {
		problems[i] = Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func startTimer(timer *time.Timer, numberOfQuestions int, score *int) {
	<-timer.C
	fmt.Printf("\nTime is up! You score %d out of %d\n", *score, numberOfQuestions)
	os.Exit(0)
}

func main() {
	config := parseConfig()
	lines := readFile(config.csv)
	problems := parseQuestions(lines)
	numberOfQuestions := len(problems)

	score := 0

	timer := time.NewTimer(time.Duration(config.limit) * time.Second)
	go startTimer(timer, numberOfQuestions, &score)

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			score++
		}
	}

	timer.Stop()

	fmt.Printf("You score %d out of %d\n", score, numberOfQuestions)
}
