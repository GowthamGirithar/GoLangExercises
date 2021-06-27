package main

import (
	"bufio"
	"os"

)

//take input from user
func main() {
	reader :=bufio.NewReader(os.Stdin)
	println("Enter the string")
	data,_  := reader.ReadString('\n')   //deleimiter for the input data
	
	println("The input string is ", data)

	// space is the delimiter
	reader2 :=bufio.NewReader(os.Stdin)
	println("Enter the string")
	data2,_  := reader2.ReadString(' ')   //deleimiter for the input data
	println("The input string is ", data2)
	
	var input int
	fmt.Scan(&input)
	fmt.Print(input)

}
