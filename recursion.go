package main

import "fmt"

//This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {

	if n == 0 {
		return 1

	}

	multiply := fact(n - 1)

	sum := n * multiply
	//	fmt.Println("original ", n)

	fmt.Println("multi ", multiply)
	//fmt.Println("sum", sum)
	fmt.Println()
	return sum
}
func main() {
	result := fact(7)
	fmt.Println(result)
}
