// we use mutex when race contion occurs, when the race condition occurs? -> Two or more goroutines trys access the same variable at the same time and at least one of them modifies it.
//how's mutex work?? like this ->
//🚶
// Lock Door
// ↓
// Bathroom //only one persone useot as a time, same only one goroutine use some resourse at a time.
// ↓
// Unlock Door
// ↓
//Next Person

package main

import (
	"fmt"
	"sync"
)


var count int
var mu sync.Mutex
func increment(wg* sync.WaitGroup){
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock() //use unlock like this, why? ->  The program panics before Unlock() is reached, means any error!!. Now the mutex stays locked forever.
	count++
	//mu.Unlock()
}

func main(){
var wg sync.WaitGroup
for i:=1; i<100; i++{
	wg.Add(1)
	go increment(&wg)
}
wg.Wait()
fmt.Println("done Bro....")
}

// Count
// ↓
// Lock
// ↓
// G1
// ↓
// Unlock
// ↓
// Lock
// ↓
// G2
// ↓
// Unlock
// ↓
// Lock
// ↓
// G3