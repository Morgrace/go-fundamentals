package main

import "fmt"

var subjectOffered uint
var grades = []uint{}

func main() {
	fmt.Println("How many subject did you take ?")
	fmt.Scan(&subjectOffered)

	// Storing grades per subject
	for i := 0; i < int(subjectOffered); i++ {
		var gradeInput int
		fmt.Printf("Enter grade for subject %v:\n", i+1)
		fmt.Scan(&gradeInput)

		grades = append(grades, uint(gradeInput))
	}

	// Calculating sum of the grades
	var sum uint
	for _, grade := range grades {
		sum += grade
	}

	average := int(sum) / len(grades)
	fmt.Println("======== Results ========")
	fmt.Printf("Your grades: %v\n", grades)
	fmt.Printf("Average: %v\n", average)

}
