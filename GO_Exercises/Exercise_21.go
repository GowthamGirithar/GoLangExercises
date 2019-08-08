package main
//channel - part 2
func main() {
	data := make(chan int)
	close(data) //closing of the channel

	_, ok := <- data   // ok will false if channel is closed and no more read operations can be performed
	println(ok)
}

func main_1() {

	// in this example we can see one go routine is reading the data
	// main go routine is writing the data
	// if we comment the reading part , it will throw error because no one is there to read it
	data := make(chan int)
	go readDatas(data)
	data <- 20
}

func readDatas(data chan int){
	val, ok := <- data   // ok will false if channel is closed and no more read operations can be performed
	if(ok){
		println(val)
	}



}


//Just to help you understand blocking concept, first send operation data <- 30 is blocking
// and some goroutine has to read data from the channel, hence readDatas goroutine is scheduled by the Go Scheduler.

// if you try to read again when the data is not there ,if there is Second read operation that will be blocking because channel c does not have any data to be read from,
// hence Go Scheduler activates main goroutine and program starts execution from close(c) function.