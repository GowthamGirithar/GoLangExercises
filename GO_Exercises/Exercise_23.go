package main

func main() {
	data := make(chan int)
	go print(data)

	for c:= range data{
		println(c)
	}

	//in this case no need of checking the channel is closed or not
}

func print(data chan int) {

	for i:=0 ; i<=10;i++{
		data <- i
	}

	close(data) //A channel can be closed so that no more data can be sent through it.
}

