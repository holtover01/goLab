package main

import "fmt"

//Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	//Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a + b
}

//When you have multiple consecutive parameters of the same type,
//you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

//now we can pass however many int values we want
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2, 3)
	sum(1)
	nums := []int{1, 2, 3, 4}
	sum(nums...) //passes the entire nums values

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals()
	fmt.Println(c)
	//Call a function just as you’d expect, with name(args).
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
