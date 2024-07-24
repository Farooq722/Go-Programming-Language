package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to mytime")

	todayDate := time.Now()

	fmt.Println("today date is :", todayDate)

	fmt.Println(todayDate.Format("02-01-2006 Monday"))
}
