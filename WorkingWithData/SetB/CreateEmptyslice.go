package main

import "fmt"

func main() {
    // Method 1: Using a slice literal
    var emptySlice1 []int
    fmt.Println("Empty Slice 1 (using slice literal):", emptySlice1, "Length:", len(emptySlice1), "Capacity:", cap(emptySlice1))

    // Method 2: Using the make function
    emptySlice2 := make([]int, 0)
    fmt.Println("Empty Slice 2 (using make):", emptySlice2, "Length:", len(emptySlice2), "Capacity:", cap(emptySlice2))

    // Method 3: Declaring and appending elements later
    emptySlice3 := []int{}
    fmt.Println("Empty Slice 3 (declared and initialized):", emptySlice3, "Length:", len(emptySlice3), "Capacity:", cap(emptySlice3))

    // Appending elements to an empty slice
    emptySlice1 = append(emptySlice1, 10, 20, 30)
    fmt.Println("\nAfter appending elements to Empty Slice 1:", emptySlice1, "Length:", len(emptySlice1), "Capacity:", cap(emptySlice1))

    // Checking if a slice is nil
    if emptySlice1 == nil {
        fmt.Println("Empty Slice 1 is nil")
    } else {
        fmt.Println("Empty Slice 1 is not nil")
    }

    if emptySlice2 == nil {
        fmt.Println("
