package main

import "fmt"

func main() {
	// Declare and initialize a slice
	slice1 := []int{1, 2, 3, 4, 5}

	// Create an empty slice and dynamically append elements
	var slice2 []int
	slice2 = append(slice2, 10)
	slice2 = append(slice2, 20, 30, 40)

	// Print the first slice
	fmt.Println("Slice1 with predefined values:")
	for i, v := range slice1 {
		fmt.Printf("slice1[%d] = %d\n", i, v)
	}

	// Print the dynamically built slice
	fmt.Println("\nSlice2 with dynamically appended values:")
	for i, v := range slice2 {
		fmt.Printf("slice2[%d] = %d\n", i, v)
	}

	// Demonstrate slice capacity and length
	fmt.Printf("\nSlice1 Length: %d, Capacity: %d\n", len(slice1), cap(slice1))
	fmt.Printf("Slice2 Length: %d, Capacity: %d\n", len(slice2), cap(slice2))
}
