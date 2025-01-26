// 6. Variadic function
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
}
