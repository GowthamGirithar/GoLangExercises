package main

func main() {
	data := make(chan int)
	go printFor(data)

	for {
		val, ok := <- data
		if(ok){  // channel is open or read can be performed
			println(val ," " ,ok)
		}else{ // channel is closed and no more read data -> and condition
			println(val , " ",ok)   // value will be default value of that data type
			break;
		}
	}
}

func printFor(data chan int) {

	for i:=0 ; i<=10;i++{
		data <- i
	}

	close(data) //A channel can be closed so that no more data can be sent through it.
}
