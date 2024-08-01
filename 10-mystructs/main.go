package main

import(
	"fmt"
)

func main() {
	fmt.Println("Welcome to STRUCTS in GoLang")

	fara := User{"farooq", 21 ,"farooq@gmail.com", "adoni"}

	fmt.Println(fara)
	fmt.Printf("User info is : %+v\n",fara)
	fmt.Println("User Name is %v\n",fara.Name)

}

type User struct{
	Name string
	Age int 
	Email string
	Address string 
}