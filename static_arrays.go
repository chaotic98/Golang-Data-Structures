package main

import "fmt"

func main() {
	// Declare an array with explicit length and default zero values
	var arr1 [5]int

	// Declare and initialize an array with specific values
	arr2 := [5]int{1, 2, 3, 4, 5}

	// Iterate over the first array and print its elements
	fmt.Println("Array with explicit length and default values:")
	for i, v := range arr1 {
		fmt.Printf("arr1[%d] = %d\n", i, v)
	}

	// Iterate over the second array and print its elements
	fmt.Println("\nArray with initialization values:")
	for i, v := range arr2 {
		fmt.Printf("arr2[%d] = %d\n", i, v)
	}
}
