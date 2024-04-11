//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
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

type Rectangle struct {
	length float64
	width  float64
}

func calculateArea(r Rectangle) float64 {
	return r.length * r.width
}

func calculatePerimeter(r Rectangle) float64 {
	return 2 * (r.length + r.width)
}

func main() {
	board := Rectangle{5, 10}
	tile := Rectangle{length: 20, width: 64}

	fmt.Println(calculatePerimeter(board))
	fmt.Println(calculateArea(board))

	fmt.Println(calculatePerimeter(tile))
	fmt.Println(calculateArea(tile))

	doubleBoard := Rectangle{board.length * 2, board.width * 2}

	fmt.Println(calculatePerimeter(doubleBoard))
	fmt.Println(calculateArea(doubleBoard))
}
