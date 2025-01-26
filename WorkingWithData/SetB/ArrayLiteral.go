package main

import "fmt"

func main() {
	// Initialize a single-dimensional array using an array literal
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Single-dimensional array:")
	fmt.Println(numbers)

	// Initialize a multi-dimensional array using an array literal
	matrix := [2][3]int{
		{1, 2, 3}, // First row
		{4, 5, 6}, // Second row
	}
	fmt.Println("\nMulti-dimensional array:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// Using an ellipsis (...) to let Go infer the size of the array
	inferredArray := [...]string{"Go", "is", "fun"}
	fmt.Println("\nArray with inferred size:")
	fmt.Println(inferredArray)
}
