package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const LINE_SEPARATOR = "\n"
const QUESTION_SEPARATOR = ","

type Config struct {
	csv string
}

type Question struct {
	question string
	answer   string
}

func parseConfig() Config {
	csv := "problems.csv"

	args := os.Args[1:]
	numberOfArgs := len(args)

	for i := 0; i < numberOfArgs; i++ {
		arg := args[i]

		if arg == "-h" || arg == "-help" {
			println("Usage of ./quiz_v1:")
			println(" -csv string")
			println(" 	a csv file in the format of 'question,answer' (default \"problems.csv\")")
			os.Exit(0)
			break
		}

		if arg == "-csv" {
			csv = args[i+1]
		}
	}

	return Config{csv}
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

func main() {
	config := parseConfig()
	file := readFile(config.csv)
	questions := parseQuestions(file)

	score := 0

	for _, question := range questions {
		fmt.Printf("%s=", question.question)
		reader := bufio.NewReader(os.Stdin)
		answerWithNewLine, _ := reader.ReadString(LINE_SEPARATOR[0])
		answer := answerWithNewLine[0 : len(answerWithNewLine)-1]
		if answer == question.answer {
			score++
		}
	}

	numberOfQuestions := len(questions)
	fmt.Printf("You score %d out of %d\n", score, numberOfQuestions)
}
