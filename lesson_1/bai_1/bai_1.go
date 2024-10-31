package main

import (
	"fmt"
)

func main() {
	var width int
	var height int
	fmt.Print("Enter Your width: ")
	fmt.Scanf("%d", &width)
	fmt.Print("Enter Your height: ")
	fmt.Scanf("%d", &height)
	fmt.Println("Your perimeter is: ", (width+height)*2)
	fmt.Println("Your area is: ", width*height)
}