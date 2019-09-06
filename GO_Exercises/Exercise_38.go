
package main



import "log"

func f() {
	defer func() {
		//recover carries the data , so based on the error code , you can perform some operations
		if r := recover(); r != nil {
			log.Printf("recover:%#v", r) //recover test
		}
	}()
	panic("test")
	panic(2)
}

func main() {
	f()
}

