package main

import "time"

// go to default case in select after the some time
func main() {
	chan1 := make(chan int  , 1)
	chan2 := make(chan int , 1)
	go callFunc31(chan1)
	go callFunc32(chan2)


	select {
	case res:= <- chan1 :
		println("the data from channel 1 is " , res)
	case res:= <- chan2 :
		println("the data from channel 2 is " , res)
	case <- time.After(1*time.Second) :
		println("the default due to timeout")   // after 1 second it will be executed
	}

	println("the main end")
}

func callFunc31(data chan int) {
	time.Sleep(2 * time.Second) // to check timeout scenario
	data <- 30
}

func callFunc32(data chan int) {
	time.Sleep(2 * time.Second)// to check timeout scenario
	data <- 20
}