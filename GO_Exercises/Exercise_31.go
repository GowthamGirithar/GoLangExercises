package main
// empty select -> leads to deadlock
func main() {
	
	data := make(chan int)
	
	data <- 30
	go emptySelect(data)
	select {}

	// since case statement is required to unblock or default case is required
	// here empty select is used , so total program got into deadlock mode
	
}

func emptySelect(data chan int) {
	<- data
}
