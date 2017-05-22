package main

import "fmt"

//this is not a pointer
func zeroval(ival int) {
	ival = 0
}

//this is a pointer
func zeroptr(iptr *int) {
	*iptr = 0
}
func main() {
	i := 1
	fmt.Println("initial:", i) //prints 1
	zeroval(i)
	fmt.Println("zeroval:", i) //prints 1

	//The &i syntax gives the memory address of i, i.e. a pointer to i.
	zeroptr(&i)                //passes memory address
	fmt.Println("zeroptr:", i) //0

	//Pointers can be printed too.
	fmt.Println("pointer:", &i) //prints the memory id
}
