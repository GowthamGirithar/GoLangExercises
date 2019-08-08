package main

import (
	"fmt"
)

func  main()  {
	arrayDemo();
	arrayIteration();
}

func arrayIteration() {
	fmt.Println("Array Iteration")
	var arr= [3] int{1, 2, 3}
	i:=0;
	for i < len(arr){
		fmt.Println(arr[i])
		i++;
	}

	
}

func  arrayDemo() {
	/*size of the array is important and  it is of fixed size*/
	/*when we copy it will do deep copy*/
	var arr= [3] int{1, 2, 3}
	fmt.Println(arr[0])
	var b= arr;
	b[0] = 10;
	fmt.Println("After copying and changing the value")
	fmt.Println(arr[0])
	fmt.Println(b[0])

	var c = &arr;
	c[0] = 10;
	fmt.Println("After copying and changing the value  with references")
	fmt.Println(arr[0])
	fmt.Println(c[0])
}