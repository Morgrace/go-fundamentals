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

var studentScorePerSubject = make(map[string]int)
var subject string
var score int

func gradeCalculator() {
	numberOfSubjects := 1
	maxSubject := 9
	for numberOfSubjects <= maxSubject {

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

func main() {
	gradeCalculator()
}
