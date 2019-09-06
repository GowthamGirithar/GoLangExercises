package main

import "fmt"

//constant at packege level
// public means , constant name starts with capital
// for only pacake means , starts with small case
const value = "hello"

//d := 10 - compile error says we dont knpw what is the d , is it type or const or what, so we have to use as var d =10

func main() {
	var a, b int = 10, 50
	c := 40
	var d float32 = 20.0
	isAvailable := true
	//constant at function level
	const value = "hello"

	fmt.Println("hello the value of a and b is ", a+c+b, int(d)+c, isAvailable)
	if isAvailable {
		fmt.Println("True Block")
	} else {
		fmt.Println("False Block")
	}
	/**for loop*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	/**while loop - If you skip the init and post statements, you get a while loop.*/
	for a < 100 {
		fmt.Println(a)
		a++
	}

	/**infinite loop - if we skip condition*/

	/**break with label*/
	for k := 0; k < 10; k++ {
	Test:
		for i := 0; i < 10; i++ {
			for j := 0; ; j++ {

				if j > 5 {

					break Test
				}
			}
		}

	}

}
