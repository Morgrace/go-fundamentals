package main

import (
	"fmt"
)

// 1. Student Grade Calculator (Easy-Medium)
// Build a program that:

// Takes multiple student scores as input
// Calculates average, highest, and lowest scores
// Assigns letter grades (A, B, C, D, F) based on score ranges
// Uses switch statement for grade assignment
// Allows adding multiple students until user exits

const MAX_SUBJECT = 9

var studentScorePerSubject = make(map[string]int)
var subject string
var score int

func gradeCalculator() {
	numberOfSubjects := 1
	for numberOfSubjects <= MAX_SUBJECT {

		fmt.Printf("\n>>Please enter subject (%d)\n>>", numberOfSubjects)

		fmt.Scan(&subject)

		fmt.Println("\nPlease enter score🎯")
		fmt.Scan(&score)

		storeScorePerSubject(subject, score)
		numberOfSubjects++
	}

	scores := make([]int, len(studentScorePerSubject))
	for _, score := range studentScorePerSubject {
		scores = append(scores, score)
	}
	averageScore := calcAverage(scores)
	highestScore := findMax(scores)
	lowestScore := findMin(scores)

	fmt.Println("HIGHEST SCORE:", highestScore)
	fmt.Println("LOWEST SCORE:", lowestScore)
	fmt.Println("AVERAGE SCORE:", averageScore)

	fmt.Printf("BELOW ARE YOU GRADES PER SUBJECT:\n%v\n", gradeAssignment(studentScorePerSubject))
}

func gradeAssignment(data map[string]int) map[string]string {
	subjectGrade := map[string]string{}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	for key, value := range data {

		if value < 0 || value > 100 {
			errorMessage := fmt.Sprintf("\nUnknown Error Has occurred: Invalid score %v\n", value)
			panic(errorMessage)
		}

		switch {
		case value >= 80:
			subjectGrade[key] = "A"
		case value >= 65:
			subjectGrade[key] = "B"
		case value >= 50:
			subjectGrade[key] = "C"
		case value >= 40:
			subjectGrade[key] = "D"
		case value >= 30:
			subjectGrade[key] = "E"
		default:
			subjectGrade[key] = "F"
		}
	}
	return subjectGrade
}

func storeScorePerSubject(subject string, score int) {
	studentScorePerSubject[subject] = score
	fmt.Println(studentScorePerSubject)
}

func calcAverage(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	average := float64(sum) / float64(len(numbers))
	return average
}

func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // or panic/error for empty slice
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func findMin(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // or panic/error for empty slice
	}

	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func init() {
	fmt.Println("STUDENT GRADE CALCULATOR")
}

func main() {
	gradeCalculator()
}
