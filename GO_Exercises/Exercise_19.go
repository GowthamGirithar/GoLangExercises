package main

import (
	"sync"
)

var val=0
// Read Write Mutex - this example shows how to have seperate read write lock
// nothing else- just use how to use syntax not for identical example
func main() {

	var mutex = sync.RWMutex{}
	var waitGroup =sync.WaitGroup{}

	for i :=0 ; i <10 ; i++{
		waitGroup.Add(2)
		go writeData(&mutex, &waitGroup)
		go readData(&mutex, &waitGroup)
	}

	waitGroup.Wait()
}

func readData(mutex *sync.RWMutex, group *sync.WaitGroup) {
		//no need of lock.. but still for demo how to use
		mutex.RLock()
		println(val)
		mutex.RUnlock()
	    group.Done()
}
func writeData(mutex *sync.RWMutex, group *sync.WaitGroup) {
	//no need of lock.. but still for demo how to use
	mutex.Lock()
	val++
	mutex.Unlock()
	group.Done()
}
