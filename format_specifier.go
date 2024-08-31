package main

import "fmt"

func main() {
	age := 20
	name := "Mridul"
	Gender := 'M'

	// Println statement : means print the statement in next line.
	fmt.Println("Age: ", age)
	fmt.Println("Name:", name)

	// Printf : used for fomatting
	fmt.Printf("Age: %d\n", age)       // There are many fomat specifiers in Go some are %d, %v, %s, %_, %f, %T
	fmt.Printf("Name: %s\n", name)     // %s is for strings
	fmt.Printf("Age: %f\n", age)       // %f is for float but you see here that i formating int to float that will provide you some other output
	fmt.Printf("Age: %T\n", age)       // %T is for type
	fmt.Printf("Age: %v\n", age)       // %v is for value
	fmt.Printf("Age: %x\n", age)       // %x is for hexadecimal
	fmt.Printf("Age: %b\n", age)       // %b is for binary
	fmt.Printf("Gender: %c\n", Gender) // %c is for character
	fmt.Printf("Age: %p\n", &age)      // %p is for memory address

	k := 2.5
	fmt.Printf("Value of k: %f\n", k)
	fmt.Printf("Value of k %0.2f\n", k) // Here you will notice that output we are getting is 2 decimal place only

}
