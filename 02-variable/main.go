package main

import "fmt"

//  %T is a format specifier used to find the TYPE of variable

func main(){
	//String type
	var username string = "Farooq"
	 
	fmt.Println(username)
	fmt.Printf("Variable type is  %T \n", username)

	//Bool type
	var isLogging bool = false	 

	fmt.Println(isLogging)
	fmt.Printf("Variable type is  %T \n", isLogging)

	//Integer type
	var smallVal int = 300	 

	fmt.Println(smallVal)
	fmt.Printf("Variable type is  %T \n", smallVal)

	//Float32 and Float64
	var smallFloat float32 = 300.5416574	 

	fmt.Println(smallFloat)
	fmt.Printf("Variable type is  %T \n", smallFloat)

	//Implicit type
	var myName = "Farooq"
	fmt.Println(myName)
	fmt.Printf("Variable type is  %T \n", myName)

	//No var type
	age := 21
	fmt.Println(age)
	fmt.Printf("%T",age)
}