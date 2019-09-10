package main

import (
	"fmt"
	"sync"
)
//demo on sync packages
func main() {
	var m sync.Map // it can be any key and  value -> type interface type
	_, isPresent := m.LoadOrStore("Test", 1) // check if present and return that if so, otherwise insert
	if isPresent {
		println("is Present")
	}
	data, _ := m.LoadOrStore("Test2", "te")
	fmt.Printf("%v", data)
	m.LoadOrStore(1, "te")

	m.Delete(1)
	

}
