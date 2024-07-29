package main

import(
	"fmt"
)
func main(){
	fmt.Println("Welcome to arrays")

	var arraysList [4]string

	arraysList[0] = "farooq"
	arraysList[1] = "fara"
	arraysList[2] = "aayan"
	arraysList[3] = "abdul"

	fmt.Println("arrayList is : ", arraysList)
	fmt.Println("arrayList is : ", len(arraysList))

	var arraysInt [4]int
	arraysInt[0] = 1
	arraysInt[1] = 2
	arraysInt[2] = 3
	arraysInt[3] = 4

	fmt.Println("arrays in int :", arraysInt)
	
}