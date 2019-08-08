package main

/** Errors demo*/
import (
	"fmt"
)
/** you have to use err.Error() to get the message of the error thrown*/
func main() {
	num1 := 10
	num2 := 0
	result , err :=divide(num1 , num2)
	if err != nil{
		println("Error happened" , err.Error())
	}
	println(result)

}

func divide(num1 , num2 int)  (int , error){
	if num2 == 0 {
		return 0 , fmt.Errorf(" num2 cannot be 0")
	}

	return num1/num2, nil
}

