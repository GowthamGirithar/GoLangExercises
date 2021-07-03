package main

/**functions demo*/
func main() {
	num1 := 1
	num2 := 2

	/**normally goland passes the value - but pointer helps to pass the address to work on */
	println(addReturnAddress(&num1,&num2))
	println("the value changed after the return " , num1 , num2)
	println(addChangeDateReturnAddress(&num1,&num2))
	println("the value changed after the return " , num1 , num2)
	
	var data *int= &num1 //*int means we are assigning the address variable
	println(data, *data) // address and value of 10


}



/**return the address *int means pointer variable*/
/* here we are taking the input as address and then take the value  - return the address*/
/* in other languages function block over the memory in stack is cleared. but in golang it is moved to heap memory (result address)*/
//here we are saying *int , it means we are receiving the pointer variable - it doesnt mean that we will have value and we can directly use it
func addReturnAddress( num1 , num2 *int) *int{
	*num1 = 100
	*num2 = 100
	result := *num1+ *num2
	return &result
}

/**return the address *int means pointer variable*/
/* here we are taking the input as address and then take the value  - return the address*/
/* in other languages function block over the memory in stack is cleared. but in golang it is moved to heap memory (result address)*/
func addChangeDateReturnAddress( num1 , num2 *int) *int{
	*num1 = 10
	*num2 = 10
	result := *num1+ *num2
	return &result
}
