package main

import "strings"

//string builder - it is same like java
//string += new operator will create the new string every time and string is immutable
//but string builder is efficient
func main() {

	dataStr := "the data is good"
	words := strings.Split(dataStr," ")
	var finalData string
	for c := range words{
		finalData+= words[c]
	}
	println(finalData)


	//with string builder

	var finalDataBuildr strings.Builder
	for c := range words{
		finalDataBuildr.WriteString(words[c])
	}
	println(finalDataBuildr.String())

	//we can also reset the string builder so it can be used to do other operation
	//.reset()

}
