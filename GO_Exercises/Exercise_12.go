package main

type Employee struct {
	name string
	age  int
}

type count int

/**methods */
func main() {
	//you have initialize to call the method
	emp := Employee{}
	emp.name="Gowtham"
	emp.printData()
	cnt := new(count)
	cnt.printData(10)

/**same method name but different arguement*/

}


//why method?
//1. we cannot say Go is object oriented
//but it has the methods to make happy for developers

// difference between func and methods
// methods are bound to the type and can have same method name

// methods cannot bind to the type of outside package

func (emp Employee) printData(){
	println("the data is " , emp.name , emp.age)
}

func (a count) printData (num int){
	println("the data is " , num)
}