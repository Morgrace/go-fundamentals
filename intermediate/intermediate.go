package main

import (
	"fmt"
)

func init() {
	fmt.Println("--------------------This is Intermediate Go----------------------")
}

// func main() {
// 	sequence := adder()
// 	sequence()
// 	sequence()
// 	sequence()
// }

// // CLOSURES
// func adder() func() int {
// 	i := 0
// 	fmt.Println("Initial value", i)

// 	return func() int {
// 		i++
// 		fmt.Println("added 1 to i = ", i)
// 		return i
// 	}
// }

//	func main() {
//		var ptr *int
//		var a int = 10
//		ptr = &a // referencing a pointer
//		fmt.Println(ptr)
//		fmt.Println(a)
//		fmt.Println(*ptr) // deferencing a pointer
//	}
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Strings and Runes
func main() {
	message := "Hello, Go"
	message1 := `Hello\nGo`
	message2 := "Hello \tGo"
	message3 := "Hello \rGo"
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)

	fmt.Println("Length of raw message variable is", len(message))
	fmt.Println("The first character in message is", message[0]) // ASCII

	for index, character := range message {
		fmt.Printf("String at index %d is %c\n", index, character)
	}

	p := Person{
		firstName: "morgrace",
		age:       27,
		lastName:  "Precious",
	}
	p.incrementAgebyOne()
	fmt.Println(p.fullName())
	fmt.Printf("Your name is %s and you are %d years old\n", p.firstName, p.age)

	// anonymous structs
	user := struct {
		name     string
		language string
	}{
		name:     "Ada",
		language: "Ogoni",
	}
	fmt.Println(user)

}
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}
func (p *Person) incrementAgebyOne() {
	p.age++
}
