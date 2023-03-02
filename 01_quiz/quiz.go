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
	help     bool
	filename string
}

type Question struct {
	question string
	answer   string
}

func parseConfig() Config {
	help := false
	filename := "problems.csv"

	args := os.Args[1:]
	numberOfArgs := len(args)

	for i := 0; i < numberOfArgs; i++ {
		arg := args[i]

		if arg == "-h" || arg == "--help" {
			help = true
			break
		}

		if arg == "-f" || arg == "--filename" {
			filename = args[i+1]
			continue
		}
	}

	return Config{help, filename}
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

const DEBUG = false

func main() {
	config := parseConfig()
	if DEBUG {
		fmt.Println(config)
	}

	file := readFile(config.filename)
	if DEBUG {
		fmt.Println(file)
	}

	questions := parseQuestions(file)
	if DEBUG {
		fmt.Println(questions)
	}

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
