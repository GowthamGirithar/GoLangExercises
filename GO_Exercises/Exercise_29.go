package main

import "time"

//deadlock case switch - to avoid normally default is used when no channels are available to send or receive data.
//Go does not schedule any other goroutines to send data to channels if data is not immediately available.
func main() {
	chan1 := make(chan int  , 1)
	chan2 := make(chan int , 1)
	go callFunc21(chan1)
	go callFunc22(chan2)


	select {
	case res:= <- chan1 :
		println("the data from channel 1 is " , res)
	case res:= <- chan2 :
		println("the data from channel 2 is " , res)
	default:
		println("the default ")   //nil channel will go to default case
	}

	println("the main end")
}

func callFunc21(data chan int) {
	time.Sleep(2 * time.Second)
	data <- 30
}

func callFunc22(data chan int) {
	time.Sleep(2 * time.Second)
	data <- 20
}
