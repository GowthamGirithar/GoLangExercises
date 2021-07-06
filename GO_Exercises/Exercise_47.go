package main

import "fmt"

func main() {
	//RemoveGivenElementWithoutAlteringOrder - remove 20
	var arr = [...]int{1, 20, 5, 0}
	index:=0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 20{
			index=i
			break
		}
	}
	fmt.Println(arr)//[1 20 5 0]
	arr1:= append(arr[:index], arr[index+1:]...)  // till location is there it wont create new reference
	fmt.Println(arr1)//[1 5 0]
	fmt.Println(arr)//[1 5 0 0]
	arr1=append(arr1, 10)
	fmt.Println(arr1)//[1 5 0 10]
	fmt.Println(arr)//[1 5 0 10]
	arr1=append(arr1, 20)
	fmt.Println(arr1) ////[1 5 0 10 20]
	fmt.Println(arr) //[1 5 0 10]
}
