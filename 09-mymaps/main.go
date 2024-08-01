package main

import (
	"fmt"
)
func main() {
	fmt.Println("Welcome to Maps in Golang")
	languages := make(map[string]int)

	languages["Go"] = 5
	languages["py"] = 2
	languages["c++"] = 5
	languages["js"] = 4

	fmt.Println("Mapped languages are :", languages)
	fmt.Println("Mapped language is :", languages["js"])

	for key, value := range languages{
		fmt.Println("For key %v, value is %v \n", key, value)
	}


}