package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number to display its multiplication table: ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
