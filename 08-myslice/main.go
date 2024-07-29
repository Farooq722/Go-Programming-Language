package main

import(
	"fmt"
	"sort"
)
func main(){
	fmt.Println("Welcome to Slice's ")
	
	//one of the way declaring slicing
	var sliceList = []string{"farooq"}
	fmt.Println(sliceList)

	sliceList = append(sliceList, "fara", "abdul")
	fmt.Println(sliceList)

	sliceList = append(sliceList[:2])
	fmt.Println(sliceList)

	//another way of declaring slice using make builtin function

	//1st arg is type of elements in slice and 2nd arg is len slice but we can dynamically allocate elements using append builtin function
	sliceMakeList := make([]int, 4)

	//bydefault slice fills with 0's if we don't initialize
	fmt.Println(sliceMakeList)

	sliceMakeList[0] = 124
	sliceMakeList[1] = 464
	sliceMakeList[2] = 987
	sliceMakeList[3] = 134
	fmt.Println(sliceMakeList)

	//adding elements dynamically using append 
	sliceMakeList = append(sliceMakeList,4899,566,7899)
	fmt.Println(sliceMakeList)

	//sorting
	sort.Ints(sliceMakeList)
	fmt.Println(sliceMakeList)

}