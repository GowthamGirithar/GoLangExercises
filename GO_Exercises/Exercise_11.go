package main

import "fmt"
/**functions demo*/
func main() {
	num1 := 1
	num2 := 2
	println(add(num1,num2))
	println(addReturn(num1,num2))
	/**normally goland passes the value - but pointer helps to pass the address to work on */
	println(addReturnAddress(&num1,&num2))
	println(addChangeDateReturnAddress(&num1,&num2))
	println("the value changed after the return " , num1 , num2)
	addNumbers(1,2,3,4)
	sum, err := addNumError(0 , 0)
	if err != nil {
		println(err)
	}else{
		println(sum)
	}
	/**function is assigned to variable*/
	var fun func( a int)= func( a int) {
		println("the function is called"  , a)
	}
	fun(10)

	/**function which returns the value*/
	var fun2 func( a int) int= func( a int) int {
		println("the function is called"  , a)
		return a
	}
	println(fun2(10))
}

func addNumError(i int, i2 int) (int , error) {
	if i == 0 && i2 ==0 {
		return 0,fmt.Errorf("Both Number cannot be 0")
	}
	return i+i2,nil
}
/**var args */
func addNumbers(numbers ...int) {
	sum :=0
	//note it is an array so index based operations
	for i := range  numbers {
		sum +=numbers[i]
	}
	println("the sum is " , sum)
}

/**function to add two numbers and return the sum*/
func add(num1 , num2 int) int{
	return num1+num2
}

/**function to add two numbers and void*/
func addNoReturn(num1 , num2 int) {
	println( num1+num2)
}

/**function to add two numbers and return the result data while returning - not used often*/
func addReturn(num1 , num2 int)  (result int) {
	result = num1+num2
	return
}

/**return the address *int means pointer variable*/
/* here we are taking the input as address and then take the value  - return the address*/
/* in other languages function block over the memory in stack is cleared. but in golang it is moved to heap memory (result address)*/
func addReturnAddress( num1 , num2 *int) *int{
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