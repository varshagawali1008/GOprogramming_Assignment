package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Multidimensional Slice:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
