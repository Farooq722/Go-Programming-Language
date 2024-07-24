package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Pointers")

	// var ptr *int ----- just declared the variable
	//fmt.Println("what is inside the ptr", ptr)  Nill

	number := 20
	myPtr := &number 

	fmt.Println("lets print myPtr what it holds ", myPtr)
	fmt.Println("lets print myPtr with * what it holds ", *myPtr)

	*myPtr = *myPtr + 10
	fmt.Println("number value is changed ", number)
}
