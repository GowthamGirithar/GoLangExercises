package main

import (
	"fmt"
	"strconv"
)
func main()  {
	stringDemo()
	convertString()
}

func convertString() {
	//var str string ="1"
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)
	print(b ,f , i , u , err)

	//The most common numeric conversions are Atoi (string to int) and Itoa (int to string). strconv.Atoi("-42")

}

func stringDemo() {
	var data string;
	data="madsam"
	println(data)
	println(len(data))
	isPalindrome := true;
	for i:=0 ; i< len(data)/2; i++ {
		if(data[i] != data[len(data)-i-1]){
			isPalindrome=false;
		}
	}
	println(isPalindrome)
	for i, ch := range "日本語" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}

}
