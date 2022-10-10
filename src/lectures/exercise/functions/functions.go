//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that returns any number
func anyNumber(number int) int {
	return number
}
//* Write a function that returns any two numbers
func anyTwoNumbers()(int, int) {
	return 2,4
}

//* Add three numbers together using any combination of the existing functions.
//* Write a function to add 3 numbers together, supplied as arguments, and

func add(lhs, mid, rhs int) int {
	return lhs + mid + rhs
}

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greeting(name string){
	 fmt.Println("Hello, ", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func echo()string{
	return "Hi, there!"
}

func main() {
	greeting("momo")
	fmt.Println(echo())

	fmt.Println("The answer is: ", add(1,2,3))

	first, second := anyTwoNumbers()
	fmt.Println("Return any two numbers, ", first, " and ", second)
//* Add three numbers together using any combination of the existing functions.
//  * Print the result

	fmt.Println("Add three numbers together, ",add(first, second, 3))
}
