package main

//channel
func main() {

	var data1 chan int;
	println(data1) // in this case it is 0 and nil channel i cannot pass to go routines
	//so only we have to use make which will have the address- ready to use channel
	//one go routine read from channel and other one write from channel

	//channels are default pointers
	data := make(chan int)
	println(data)  // print address
	go updateData(data) // go routine update the data
	// to read the data or write the data <- arrow is used
	msg := <- data // main go routine read the data
	println(msg)


	// when you write the data to channel , it will be blocked and we need to read by the go routines
	//consider code like below
	// data := make(chan int)
	// data <- 20
	// after this no one do read from channel , it is blocked
	// same go routine will not read the data and write it - here main go routine only
	// so it error as all are sleeping
	//channel search other go routine to schedule ..if it dont find throws the error

}


func updateData(data chan int) {
	data <- 30
}
