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

//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
type product struct {
	name string
	price int
}

func printInfo (p [4]product) {
	var	total, totalItem int
	for _, product := range p {
		total += product.price	
		if product.name != "" {
			totalItem += 1
		}
	}
	fmt.Println("total cost of the items is: ", total)
	fmt.Println("total number of items is: ", totalItem)
	fmt.Println("the last item is: ",p[totalItem - 1])
}
func main() {
  shoppingList := [4]product {
	 {"milky", 1},
	 {"soap", 2},
	 {"water", 2},
}
printInfo(shoppingList)
shoppingList[3] = product{"beef", 10}
printInfo(shoppingList)
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

}
