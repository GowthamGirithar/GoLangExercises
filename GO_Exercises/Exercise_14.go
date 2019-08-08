package main

import "fmt"

type EmployeeDto struct {
	name string
	age int
}
type EmptyInterface interface {

}
//common interface
type EmployeeService interface {
	create()
	update()
	delete()
}

// interface with one method is normally name ends with er

type Caller interface {
	call()
}

// embedded interface

type ExerciseInterface interface {
	Caller
	EmployeeService
}

/**interface demo*/
func main() {
	emp := EmployeeDto{}
	emp.name="Gowtham"
	emp.create()
	/** it is used to get all the data type and based on data type perform an operation*/
	var value interface{}
	value=90
	getType(value)

	//FUNCTION
	emptyInter(&value)

	//



}

func getType(value interface{})  {
	fmt.Printf("%T", value)
}

func (dto EmployeeDto) create()  {
	println(dto.name)
}

/**empty interface*/
func  emptyInter(value *interface{})  {
	println()
	println("the empty interface")
	println(*value)
}


