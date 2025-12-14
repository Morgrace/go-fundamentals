package main

import (
	"fmt"
	"slices"
)

const pi = 3.14
const GRAVITY = 9.81

func lessons() {
	// CONSTANTS
	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday  int = 4
	)

	// ARITHMETRIC OPERATIONS
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition", result)

	result = a - b
	fmt.Println("Substraction", result)

	result = a * b
	fmt.Println("Multiplication", result)

	result = a / b
	fmt.Println("Divison", result)

	result = a % b
	fmt.Println("Remainder", result)

	// overflow and underflow
	var maxInt int64 = 949349343493432499
	fmt.Println(maxInt)

	maxInt += 9
	fmt.Println(maxInt)

	// FOR LOOOP
	// simple iteration (counter type)
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over collection or range
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value:%d\n", index, value)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip that iteration to the next
		}
		fmt.Println("Odd Number", i)
		if i == 5 {
			break // stop the loop
		}
	}
	// New feature for loop (iterating over integers)
	for i := range 10 {
		fmt.Println(10 - i)
	}
	// using for loop like a (while loop)
	i := 1
	for i <= 5 {
		fmt.Println("Iteration", i)
		i++
	}

	// ARRAY
	// var arrayName [size]elementType
	{
		var numbers [5]int
		fmt.Println(numbers)

		// updating an array
		numbers[4] = 40
		fmt.Println(numbers)

		fruits := [4]string{"banana", "orange", "grapes", "apple"}
		fmt.Println(fruits)

		for index, value := range fruits {
			fmt.Printf("\nindex: %v\n", index)
			fmt.Printf("\nvalue: %v\n", value)
		}
	}
}
func main() {

	// SLICES
	// var sliceName[]elementType
	var numbers []int
	numbers2 := []int{9, 8, 9}

	var numbers3 = [5]int{1, 2, 3, 4, 5}
	slice := numbers3[0:4]

	slice = append(slice, 5, 6, 7)

	// copying slices
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)

	sliceCopy[0] = 50

	if slices.Equal(slice, sliceCopy) {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}

	fmt.Printf("\nEmpty slice: %v\n Slice1: %v\n slice2: %v\nAppendedSlice:%v", numbers, numbers2, slice, slice)

	fmt.Println("\ncopied slice", sliceCopy)

	// MAPS
	// var mapVarible  map[keyType]valueType

	myMap := make(map[string]int)

	myMap["key1"] = 9
	fmt.Println(myMap)

	// checks
	value, ok := myMap["key1"]
	fmt.Print(value, ok)

}
