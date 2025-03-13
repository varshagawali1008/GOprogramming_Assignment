package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Buffered channel with size 2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // Output: 1
	fmt.Println(<-ch) // Output: 2
}
