package main

// two go routine and one for square and another for cube
func main() {
	println("main start")
	squareChan := make(chan int)
	cubeChan := make(chan int)
	go square(squareChan)
	go cube(cubeChan)
	testData := 10
	squareChan <- testData
	cubeChan <- testData
	// the next line wait till the data is there
	val1 , val2 := <-squareChan,<-cubeChan
	println("the sum is ", val1+val2)
	println("main end")

	
}

func cube(data chan int) {
	value := <- data
	data <- value *value
}

func square(data chan int) {
	value := <- data
	data <- value *value * value
}
