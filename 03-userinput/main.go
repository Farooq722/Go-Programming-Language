package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	/*
	bufio.NewReader(os.Stdin) creates a new buffered reader that reads from standard input.
	fmt.Print("Enter text: ") prints a prompt for the user.
	reader.ReadString('\n') reads input from the user until a newline character is encountered.
	*/

	//UserInput
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(welcome)
	fmt.Println("Enter something")
	//comma ok syntax or err ok syntax we can either declare _ or err for error, we use _ when dont care about error if we does then writes err.
	/*
	input, _  ------- we use _ when dont care about error
	input, err ------ we use err when do care about error
								same goes for Input 
	*/
	// \n stops taking input when we hit enter button

	input, _ := reader.ReadString('\n')
	fmt.Println("my name is :", input)
	fmt.Printf("type of variable %T \n", input)
}
