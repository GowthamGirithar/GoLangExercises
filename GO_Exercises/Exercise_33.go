package main

import "strings"

//string operations  -> strings package is used
func main() {
	dataStr := "the data is big"
	//words from the string
	words := strings.Fields(dataStr)
	for c := range  words{
		println(words[c])
	}
	//comparition
	println("the string comparition" , strings.Compare("Gowtham" , "gowtham"))
	println("the string contains" , strings.ContainsAny(dataStr,"the"))
	println("the rune contains ",strings.ContainsRune(dataStr,'s'))
	// more operation you can see in strings package
}
