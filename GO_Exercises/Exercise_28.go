package main

import "time"

//select statement with unbuffered channels
func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go callFunc1(chan1)
	go callFunc2(chan2)

	// main go routine will block and it will continue when one of the case is executed
	// this is like load balancer, it wil send the request to all the available service
	// when it recieves one , it take that and continue (discard other ones)
	select {
	case res:= <- chan1 :
		 println("the data from channel 1 is " , res)
	case res:= <- chan2 :
		println("the data from channel 2 is " , res)
		/*default :
			println("the data is default ")*/
	}

	// if you enable default , it will print that one and unblock
	// so default is unblocking state
	println("the main end")
}

func callFunc2(data chan int) {
	time.Sleep(2 * time.Second)
	data <- 30
}

func callFunc1(data chan int) {
	data <- 20
}
