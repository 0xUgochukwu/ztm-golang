//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

func main() {
	shoppingList := [4]struct {
		name  string
		price float64
	}{
		{"apple", 1.99},
		{"banana", 0.99},
		{"orange", 2.99},
		{"grape", 3.99},
	}

	fmt.Println("Last item on the list:", shoppingList[len(shoppingList)-1].name)
	fmt.Println("Total number of items:", len(shoppingList))

	totalCost := 0.0
	for i := 0; i < len(shoppingList); i++ {
		totalCost += shoppingList[i].price
	}

	fmt.Println("Total cost of the items:", totalCost)
}
