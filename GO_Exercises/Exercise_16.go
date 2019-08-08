package main

import "time"
import "sync"
import "runtime"
// for demo rename main method to main one by one and test it

// difference between threads and go routines?
//Java threads map directly to OS threads, and are relatively heavyweight. Part of the reason they are heavyweight is their rather large fixed stack size. This caps the number of them you can run in a single VM due to the increasing memory overhead.
//
//Go OTOH has a segmented stack that grows as needed. They are “Green threads”, which means the Go runtime does the scheduling, not the OS. The runtime multiplexes the goroutines onto real OS threads, the number of which is controlled by GOMAXPROCS. Typically you’ll want to set this to the number of cores on your system, to maximize potential parellelism.

//go routines - it is very small compared to thread
// thread occupies more memory
func main() {

	println("THE THREAD AVAILBLE  " ,runtime.GOMAXPROCS(-1)) // IF U PASS > 0 , THEN IT WILL SET <0 , IT WILL NOT CHANGE AND RETURN THE VALUE THERE
	msg :="Hello"
	go testGO1(msg)   //it is different call stack - green threads
	msg ="Hello changed"
	println("Main call")
	time.Sleep(2 * time.Second) // for demo - making main to wait till operations are completed
}

func main_2() {

	msg :="Hello"
	go func(msg string) {
		println("test GO function " , msg)
	}(msg)
	println("Main call")
	time.Sleep(2 * time.Second) // for demo - making main to wait till operations are completed
}

func main_3() {

	// 2 GO ROUTINES PARALLEL
	go testGO1("HELLO 1")
	go testGO2("HELLO 2")
	println("Main call")
	time.Sleep(2 * time.Second) // for demo - making main to wait till operations are completed
}

var waitGroup = sync.WaitGroup{}
func main_4() {
	// instead time sleep ,use sync to wait for operatiosn to completed
	// 2 GO ROUTINES PARALLEL
	waitGroup.Add(2)   // define how many go routines
	go testGOWG("HELLO 1")
	go testGOWG("HELLO 2")
	println("Main call")
	waitGroup.Wait()// once all done complete
}


func main_5(){
 // why we have pass the data as arguement in golang?
 //when the fucntion is gonna execute , it goes to the memory of that variable and get the data

 //worng way
	for i := 0; i < 3; i++ {
		go func() {
			println(i)
		}()
	}
// correct way
	for i := 0; i < 3; i++ {
		go func(v int) {
			println(v)
		}(i)
	}

	time.Sleep(2 * time.Second)
}

func testGOWG(str string ) {
	println("the value is " , str)
	waitGroup.Done() // say once done
}




func testGO1(msg string)  {
	println("test GO 1" , msg)

}

func testGO2(msg string)  {
	println("test GO 2" , msg)
}
