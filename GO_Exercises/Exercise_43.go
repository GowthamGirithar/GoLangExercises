package main

import "fmt"

type Employee struct {
	EmpName string
}

func main(){
	emp:=getData(false)
	println(emp)
	emp =getData(true)
	fmt.Print(emp)
	emp1:=Employee{EmpName:"Gowtham"}
	//println(emp1) // ERROR , WE CANNOT PRINT THIS use fmt.Print
	fmt.Printf("the data is  ",emp1)

	//for printing , use %v for any type and for detailed print %+v
	println()
	fmt.Printf("the data is %v \n",emp1) //the data is {Gowtham}
	fmt.Printf("the data is %+v \n",emp1)//the data is {EmpName:Gowtham}

}

func getData(data bool) interface{}{
	if data{
		return Employee{EmpName:"Test"}
	}else{
		return  nil
	}

}

func getData2(data bool) Employee{
	if data {
		return Employee{EmpName: "Test"}
	} else {
		return nil // compile time error, we cannot return nil to the struct. so better to use interface return type
	}

}
