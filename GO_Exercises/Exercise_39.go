package main

import (
	"log"
	"math/rand"
	"time"
)

func main()  {
	log.SetFlags(0) //If you want to remove the formatting or change it, you can use the log.SetFlags function:
	log.Println("test")

	log.SetFlags(1)
	log.Println("test")

	rand.Seed(time.Now().UnixNano())
	for i:=0 ; i<10 ; i++{
		println(rand.Int())
	}

}
