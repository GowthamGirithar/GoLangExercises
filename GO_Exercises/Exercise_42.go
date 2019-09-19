package main

func main() {
	//new always returns the address and initialize the value with 0
	c := new([]int)
	println(c)
	println(len(*c))
	//make always returns the data and initialize the value with 0 with the capacity what we have defined
	d := make([]int, 10)
	println(d)
	println(len(d))

}

