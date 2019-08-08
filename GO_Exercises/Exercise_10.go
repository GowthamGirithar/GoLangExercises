package main

/**pointer*/
func main() {
	pointerDemo()
	pointerInternalReferDemo()
	pointerArithMeticOper()
	pointerDefault()

	//slice and map is having internal pointer and not the actual value
	//so only on assignment it is taking reference
	//this proves that assignment in golang is always a copy the value

}

func pointerDefault() {
	var a *int;
	println(a == nil)

	a= new (int)
	println(a == nil)
	println(a)
	println(*a)
}

func pointerArithMeticOper() {
	//var a= 10
	//var ptr =& a
	//ptr+1;  -> arithemetic operation in pointer is not allowed to move the pointer to different location
	//golang is for easy
}

func pointerInternalReferDemo() {
	var a= 10
	var ptr =& a
	var ptrPtr =&ptr

	println(ptrPtr)
	println(ptr)
	println(a)

	println("access the value with ptrPtr")
	println(**ptrPtr)
	**ptrPtr=11
	println(a)
}
//declare pointer
func pointerDemo() {

	var a int = 10
	//pointer is also a variable and it stored normally the address of other variable
	// we can also write as var ptr  =& a
	var ptr *int = &a

	println(a)
	println(ptr)
	// to access the value
	println(*ptr)
}
