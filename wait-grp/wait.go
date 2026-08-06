package main

import (
	"fmt"
	"sync"
)

// func hello(){
// 	fmt.Println("Hello Suru")
// }

// func main(){
// 	go hello()
// 	fmt.Println("AD")
// }
//here in above program only AD is printing after execution, why? hello func start a go routine,
//so the main func is finish immediately, When main() exits, the whole program exits, even if goroutines are still running.

//first waitGroup pgm

// func printHello(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Println("Hello")
// }
// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go printHello(&wg)
// 	//go printHello(&wg) // when i add this goroutine and add value is still 1 then the ouput print only one hello, bcs the prgm wait only for 1 goroutine, after the execution of 1st goroutine it execute the main function or amy be second goroutine if it is seduled by the go before main func.
// 	wg.Wait()
// 	fmt.Println("Main function Executed")
// }

// Main
//  │
//  │ Add(1)
//  ▼
// Counter = 1
//  │
//  │
//  ├──────────────► Goroutine
//  │                 │
//  │                 ▼
//  │            Print Hello
//  │                 │
//  │            Done()
//  │                 │
//  │                 ▼
//  │          Counter = 0
//  │
//  ▼
// Wait()
//  │
//  ▼
// Main Finished

// func task(name string, wg* sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Println(name)
// }

// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go task("task 1",&wg)
// 	go task("task 2",&wg)
// 	go task("task 3",&wg)
// 	go task("task 4",&wg)
// 	wg.Wait()
// 	fmt.Println("All the task are done.....")

// }


func printNum(num int,wg* sync.WaitGroup ){
	defer wg.Done()
	fmt.Println("Number: ", num)
}

func main(){
	var wg sync.WaitGroup
	for i:=1; i<=10; i++{
		wg.Add(1)
		go printNum(i,&wg)
	}
	wg.Wait()
	fmt.Println("Done with the task...")
}