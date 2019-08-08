package main
//buffer size - block and unblock

func main_11() {

	data := make(chan int , 2)   // here 2 is capacity of channel , if we dont provide it is 0


	// o size is called unbuffered channel- whatever written to channel is availbale to read immediately
	// size > 0 means buffered channel


	// in this case, it will not block the current thread till the size is met
	go printChannelData(data)
	data <- 20
	data <- 30
}

func main() {

	data := make(chan int , 2)   // here 2 is capacity of channel , but given 3 data

	// in this case, it will block the current thread becuase the data is more than buffer size
	go printChannelData(data)
	data <- 20
	data <- 30
	data <- 40

	println("length of channel " , len(data)) // data in queue - but it print different due to parallel communication
	println("capacity of channel " , cap(data))  // capacity given in the channel
}

func printChannelData(data chan int) {
	for c:= range data{
		println(c)
	}

	println("the data channle print")
	
}
