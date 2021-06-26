package main
// unidirectional channels
// read only and write only channels
func main_111() {
	readCh := make(<-chan int)
	writeCh := make(chan <- int)
	println(readCh )
	println(writeCh )
}


func main() {
	data := make(chan int)
	go unidirectionalChannel(data)
	data <- 20  // main go routine can read and write the data

	
}
//this go routine only can read the data
func unidirectionalChannel(data <-chan int) {
	println(<-data)
	//data <- 20   -> will throw compile time error
}
//in general, we will define the channel as normal one and in particular method we will receive it as read only or write only.
