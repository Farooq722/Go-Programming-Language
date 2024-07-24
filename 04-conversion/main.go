package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello game")

	fmt.Println("Enter the any number")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your desired number,", input) 

	//Here type of the variable " input " is a string     Before conversion
	fmt.Printf("type of variable is %T \n", input)

	//Lets see after conversion from string to int or any other type

	/*
		strconv is conversion package and we are parsing into float using parseFloat Function 
		we pass strings as a FIRST arguments and trimSpace for trimming space again in trimSpace function we pass our actual variable which we need to convert according to , AND second argument either float32 (32) or float64(64) (for this proble float is required)
	*/
	conversion, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	//we have to handle err if occurs else it will print output
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("added 1 to rating : ", conversion+1)
	}

	fmt.Printf("type of variable is %T \n", conversion)
}
