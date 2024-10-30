package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter Your String: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	if len(input)%2 == 0 {
		print("true")
	} else {
		print("false")
	}
}