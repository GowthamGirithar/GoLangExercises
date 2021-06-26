package main
// in previous example , because of channels concept
// we no need to use the time.Sleep to wait for the go routine execution - channel will block
// in this example we have seen if we dont  provide the buffer,  every send operation on channel blocks the current goroutine.
// comment the line 11 and check & uncomment and check


func main() {
	data := make(chan int)
	go getData(data)
	//data <- 30     //- when we disable this line , it will not block the current go routine
	// so main method will ovewr and go routine will run at some time
	// when we enable the data <- 30  , it will block the current go routine (main in this case) till other go routine read the data
	// -  every send operation on channel blocks the current goroutine.
	println("main over")
	//if we have any read operation in main go routine, it wil check any write is there
	// if we have any write operation in main go routine, it will check any read is there
}

func getData(data chan int) {
	<- data
	println("the data getting over")
}
