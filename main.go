package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// loadQuestions reads questions from a CSV file and sends them to questionChan.
func loadQuestions(csvPath string, questionChan chan []string) {
	file, err := os.Open(csvPath)
	if err != nil {
		fmt.Println("Error loading questions:", err)
		close(questionChan)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		questionChan <- line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading CSV:", err)
	}

	close(questionChan)
}

// readAnswers continuously reads from stdin and sends inputs to answerChan.
func readAnswers(answerChan chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer := strings.TrimSpace(scanner.Text())
		answerChan <- answer
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("\nError reading input:", err)
		close(answerChan)
		os.Exit(1)
	}
}

// askQuestion presents a question, waits for an answer or timeout, and returns the score.
func askQuestion(ans string, limit int, answerChan chan string) int {
	timer := time.NewTimer(time.Duration(limit) * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("\nTime's Up!")
		return 0
	case userAns := <-answerChan:
		if !timer.Stop() {
			<-timer.C
		}
		// Compare answers
		ansFloat, errAns := strconv.ParseFloat(ans, 64)
		userAnsFloat, errUserAns := strconv.ParseFloat(userAns, 64)

		if errAns == nil && errUserAns == nil {
			if ansFloat == userAnsFloat {
				fmt.Println("Correct!")
				return 1
			}
			fmt.Println("Incorrect!")
			return 0
		}

		if strings.EqualFold(userAns, ans) {
			fmt.Println("Correct!")
			return 1
		}
		fmt.Println("Incorrect!")
		return 0
	}
}

// singleAnswerQuestion is a wrapper for askQuestion.
func singleAnswerQuestion(question, answer string, limit int, answerChan chan string) int {
	fmt.Printf("Question: %s ", question)
	return askQuestion(answer, limit, answerChan)
}

func multipleAnswerQuestion(line []string, limit int, answerChan chan string) int {
	question, optionA, optionB, optionC, optionD, correctOption := strings.TrimSpace(line[0]), strings.TrimSpace(line[1]), strings.TrimSpace(line[2]), strings.TrimSpace(line[3]), strings.TrimSpace(line[4]), strings.TrimSpace(line[5])

	fmt.Printf("Question: %s\n", question)
	fmt.Printf("A) %s\n", optionA)
	fmt.Printf("B) %s\n", optionB)
	fmt.Printf("C) %s\n", optionC)
	fmt.Printf("D) %s\n", optionD)
	fmt.Print("Your answer (A/B/C/D): ")

	return askQuestion(correctOption, limit, answerChan)
}

// runQuiz processes each question and accumulates the score.
func runQuiz(questionChan chan []string, limit int, answerChan chan string) int {
	score := 0

	for line := range questionChan {
		if len(line) == 2 {
			score += singleAnswerQuestion(line[0], line[1], limit, answerChan)
		} else if len(line) == 6 {
			score += multipleAnswerQuestion(line, limit, answerChan)
		} else {
			fmt.Println("Invalid question format:", line)
		}
	}

	return score
}

func main() {
	// Define command-line flags
	csvPath := flag.String("csv", "problems.csv", "A csv file in the format of 'question, answer' (default 'problems.csv')")
	limit := flag.Int("limit", 10, "The time limit for each question in seconds (default 10)")
	flag.Parse()

	// Initialize channels
	questionChan := make(chan []string)
	answerChan := make(chan string, 1)

	// Start loading questions and reading answers concurrently
	go loadQuestions(*csvPath, questionChan)
	go readAnswers(answerChan)

	// Run the quiz
	score := runQuiz(questionChan, *limit, answerChan)
	fmt.Printf("\nQuiz Completed! Your Final Score is %d\n", score)
}
