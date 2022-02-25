package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
		fmt.Println(result)
	}
	return result
}
func main() {
	fmt.Println("Variadic function : Sum !!!")
	println("1+2+3+4+5 = ", sum(1, 2, 3, 4, 5))
}
