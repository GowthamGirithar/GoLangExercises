package main

import "sync"
//locking with Mutex

//we have to lock and unlock
// to make sure only one go routine is running the peice of code at a time to prevent the race condition
//


var x  = 0

// comments the unlock and lock to see how the final result changes
//because it is shared data so lock to prevent the race condition

// in this example we have used the mutex and waitgroup via pointer type becuase of local varibales

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	println("final value of x", x)
}