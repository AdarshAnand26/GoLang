//channels like a bridge, assume it like two rooms connecting usiing channels
// Room A
// |
// |  sends message
// |
// ======== Channel ========
// |
// |
// Room B

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func num(ch chan int){
// 	ch<-100 // giving the value in channels
// }

// func main(){
// 	ch:=make(chan int) // creating channel
// 	go num(ch) // calling gorutine function
// 	number:=<-ch //taking value to print
// 	fmt.Println(number)
// }

// Goroutine
//     |
//     | 100
//     V
// +-----------+
// |  Channel  |
// +-----------+
//     |
//     V
// Main

//if we dont either send or recive the value in channel then it will go to deadlock and error occurs

// func square(num int, ch chan int){
// 	ch<-num*num
// }

// func main(){
// 	ch:=make(chan int)
// 	go square(5, ch)
// 	ans:= <-ch
// 	fmt.Println(ans)
// }

func procssNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing Number: ", num)
		time.Sleep(time.Second)
	}
}
func main() {
	numChan := make(chan int)
	go procssNum(numChan)
	for {
		numChan <- rand.Intn(100)
	}

}
