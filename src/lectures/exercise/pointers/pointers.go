//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Item struct {
	name        string
	securityTag bool
}

func alterSecurityTag(item *Item) {
	item.securityTag = !item.securityTag
}

func checkout(items []Item) {
	for item := range items {
		items[item].securityTag = false
	}
}

func main() {
	items := []Item{
		{"apple", true},
		{"banana", true},
		{"orange", true},
		{"grape", true},
	}
	fmt.Println("Initial state:")
	for _, item := range items {
		fmt.Println(item)
	}
	alterSecurityTag(&items[1])
	fmt.Println("After altering security tag:")
	for _, item := range items {
		fmt.Println(item)
	}
	checkout(items)
	fmt.Println("After checkout:")
	for _, item := range items {
		fmt.Println(item)
	}

}
