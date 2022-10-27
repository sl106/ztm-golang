//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

//* Create a rectangle structure containing its coordinates
type rectangle struct {
	a coordinate //top left
	b coordinate // bottom right
}

type coordinate struct {
	x, y int
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//
func (r *rectangle) perimeter() int {
	height, length := r.size()
	return height*2 + length*2
}

func (r *rectangle) size() (int, int) {
	height := r.a.y - r.b.y
	length := r.b.x - r.a.x

	return height, length
}

func (r *rectangle) area () int {
	height, length := r.size()	
	return height * length
}

func main() {
	rec := rectangle{coordinate{0, 7}, coordinate{10, 0}}
	fmt.Printf("perimeter is: %d and area is: %d ", rec.perimeter(), rec.area())
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
	rec.a.y *= 2
	rec.b.x *= 2
	fmt.Printf("perimeter is: %d and area is: %d ", rec.perimeter(), rec.area())
}
