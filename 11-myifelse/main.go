package main

import(
	"fmt"
)

func main() {
	fmt.Println("Welcome to If Else tut")

	count := 25

	var result string

	if count < 15 {
		result = "count is less than 15" 
	} else if count > 15{
		result = "count is greater than 15"
	} else{
		result = "count is equall"
	}

	fmt.Println(result)

	if num := 5; num >= 5{
		fmt.Println("yes its greater than :", num)
	}else{
		fmt.Println("no, its not greater than  :", num)
	}

	if number:= 10; number%2 == 0 {
		fmt.Println("its even")
	} else{
		fmt.Println("its odd")
	}
}