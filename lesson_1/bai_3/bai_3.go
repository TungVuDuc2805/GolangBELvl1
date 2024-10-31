package main

import (
	"fmt"
	"sort"
)
func main() {
	n := getLen()
    array := getValues(n)
	sum := 0
	max := 0
	min := 1_000_000_000

	for i := range n {
		sum += array[i]
		if array[i] > max {
			max = array[i]
		}

		if array[i] < min {
			min = array[i]
		}
	}

	fmt.Println("Array sum: ", sum)
	fmt.Printf("Average: %.1f\n", float64(sum)/float64(n))
	fmt.Println("Maximum number: ", max)
	fmt.Println("Minimum number: ", min)
	sort.Ints(array)
	fmt.Println("Sorted array: ", array)
}

func getLen() int {
	var n int
	fmt.Println("Numbers of elements in your array: ")
	fmt.Scan(&n)
	for n <= 0 {
		fmt.Println("Please enter a positive number")
		n = getLen()
	}
	return n
}

func getValues(n int) []int {
    a := make([]int, n)
    for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i)
        fmt.Scan(&a[i])
    }

	return a
}