package main

import "os"

/**defer , panic  and recover*/
func main()  {

	deferDemo()
	deferRealDemo()
	panicDemo()
	
}

/**it is like stopping the execution ,u can recover but it will stop that method execution and follow the other ones*/
func panicDemo() {
	panic("end of panic demo method")
}

/**Here after creating , dev can immediately write close - but it will do at the end only*/
/**if err is there then panic*/
func deferRealDemo() {
	println("start of real demo of defer method")
	f, err := os.Create("/tmp/defer.txt")
	defer func() {
		f.Close()
		//after panic it execute the defer , so it is the best place if you want to recover
		recover()

	}()  //() is only invoking this anonymous function - function without name
	if err != nil {
		panic(err)
	}
	println("end of real demo of defer method")
}

//defer - it is used to say execute at the last before returing - kind of finally block
// will execute before any error panic
//it helps developer open and close immediately with defer.. otherwise developer have chance to forget
//defer works as LIFO
func deferDemo() {
	a:=10;
	println("start")
	defer println("middle1" , a)
	defer println("middle2", a)
	a=11
	println("end")
}