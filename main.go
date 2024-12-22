package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	csvPath := flag.String("csv", "problems.csv", "A csv file in the format of 'question, answer' (default 'problems.csv')")
	// limit := flag.Int("limit", 30, "The time limit for the quiz in seconds (default 30)")
	flag.Parse()

	file, err := os.Open(*csvPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count, score := 0, 0
	for scanner.Scan() {
		count++
		line := strings.Split(scanner.Text(), ",")
		question, ans := strings.TrimSpace(line[0]), strings.TrimSpace(line[1])

		fmt.Fprintf(os.Stdout, "Question %d: %s ", count, question)

		// Read user's answer
		ansScanner := bufio.NewScanner(os.Stdin)
		if !ansScanner.Scan() {
			fmt.Println("\nError reading input.")
			break
		}
		userAns := strings.TrimSpace(ansScanner.Text())

		// Try to parse both the user's answer and the actual answer as floats
		ansFloat, errAns := strconv.ParseFloat(ans, 64)
		userAnsFloat, errUserAns := strconv.ParseFloat(userAns, 64)

		if errAns == nil && errUserAns == nil {
			// Both are floats, compare numerically
			if ansFloat == userAnsFloat {
				score++
			} else {
				fmt.Printf("\nIncorrect. Game Over. Your Score is %d\n", score)
				os.Exit(0)
			}
		} else {
			// Compare as strings if parsing fails
			if userAns == ans {
				score++
			} else {
				fmt.Printf("\nIncorrect. Game Over. Your Score is %d\n", score)
				os.Exit(0)
			}
		}
	}
	fmt.Printf("\nQuiz Completed! Your Final Score is %d\n", score)
}
