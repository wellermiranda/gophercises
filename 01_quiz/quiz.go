package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const LINE_SEPARATOR = "\n"
const QUESTION_SEPARATOR = ","

type Config struct {
	csv   string
	limit int
}

type Question struct {
	question string
	answer   string
}

func parseConfig() Config {
	csv := "problems.csv"
	limit := 30

	args := os.Args[1:]
	numberOfArgs := len(args)

	for i := 0; i < numberOfArgs; i++ {
		arg := args[i]

		if arg == "-h" || arg == "-help" {
			println("Usage of ./quiz:")
			println(" -csv string")
			println(" 	a csv file in the format of 'question,answer' (default \"problems.csv\")")
			println(" -limit int")
			println(" 	the time limit for the quiz in seconds (default 30)")
			os.Exit(0)
			break
		}

		if arg == "-csv" {
			csv = args[i+1]
		}

		if arg == "-limit" {
			newLimit, err := strconv.Atoi(args[i+1])
			handleError(err)
			limit = newLimit
		}
	}

	return Config{csv, limit}
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) string {
	file, err := os.Open(filename)
	handleError(err)

	buffer := make([]byte, 500)
	data, err := file.Read(buffer)
	handleError(err)

	return string(buffer[:data])
}

func parseQuestions(file string) []Question {
	lines := strings.Split(file, LINE_SEPARATOR)
	questions := make([]Question, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, QUESTION_SEPARATOR)
		questions[i] = Question{question: parts[0], answer: parts[1]}
	}

	return questions
}

func startTimer(timer *time.Timer, numberOfQuestions int, score *int) {
	<-timer.C
	fmt.Printf("\nTime is up! You score %d out of %d\n", *score, numberOfQuestions)
	os.Exit(0)
}

func main() {
	config := parseConfig()
	file := readFile(config.csv)
	questions := parseQuestions(file)
	numberOfQuestions := len(questions)

	score := 0

	timer := time.NewTimer(time.Duration(config.limit) * time.Second)
	go startTimer(timer, numberOfQuestions, &score)

	for _, question := range questions {
		fmt.Printf("%s=", question.question)
		reader := bufio.NewReader(os.Stdin)
		answerWithNewLine, _ := reader.ReadString(LINE_SEPARATOR[0])
		answer := answerWithNewLine[0 : len(answerWithNewLine)-1]
		if answer == question.answer {
			score++
		}
	}

	timer.Stop()

	fmt.Printf("You score %d out of %d\n", score, numberOfQuestions)
}
