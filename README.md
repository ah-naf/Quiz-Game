# Quiz Game in Go

This project is a **Quiz Game** written in Go. It reads questions from a CSV file and presents them to the user one at a time. Users can provide answers within a time limit for each question. The game supports both single-answer and multiple-choice questions.

---

## Features
- **Single-Answer Questions**: Questions where the answer is a single value (e.g., `What is 2 + 2?`).
- **Multiple-Choice Questions**: Questions with options like A, B, C, and D (e.g., `Which planet is known as the Red Planet?`).
- **Time Limit**: Users must answer within the specified time; otherwise, the question times out.
- **Concurrent Input Handling**: User input is read concurrently to ensure a smooth experience.

---

## Prerequisites
- Go 1.18+ installed on your machine.
- A CSV file with questions (see format below).

---

## CSV Format
The CSV file should be structured as follows:

### Single-Answer Questions
Each line should contain a question and its correct answer:
```csv
question,answer
What is 2+2?,4
What is the capital of France?,Paris
```

### Multiple-Choice Questions
Each line should contain a question, four options (A, B, C, D), and the correct option:
```csv
question,optionA,optionB,optionC,optionD,correctOption
Which planet is known as the Red Planet?,Venus,Mars,Earth,Jupiter,B
What is the square root of 16?,2,3,4,5,C
```

---

## Usage

### Build and Run
1. Clone the repository:
   ```bash
   git clone https://github.com/ah-naf/quiz-game.git
   cd quiz-game
   ```

2. Build the project:
   ```bash
   go build -o quiz-game
   ```

3. Run the executable:
   ```bash
   ./quiz-game -csv=questions.csv -limit=15
   ```

### Command-Line Flags
- `-csv`: Path to the CSV file containing the questions (default: `problems.csv`).
- `-limit`: Time limit for each question in seconds (default: `10`).

---

## Example Output
### Single-Answer Question:
```
Question: What is 2+2? 4
Correct!
```

### Multiple-Choice Question:
```
Question: Which planet is known as the Red Planet?
A) Venus
B) Mars
C) Earth
D) Jupiter
Your answer (A/B/C/D): B
Correct!
```

### Timeout:
```
Question: What is the capital of France? 
Time's Up!
```
