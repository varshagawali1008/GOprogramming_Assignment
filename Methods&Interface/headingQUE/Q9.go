package main

import "fmt"

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func main() {
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q)
	fmt.Println("Remainder:", r)
}
