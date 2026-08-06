package main

import (
	"fmt"
	"time"
)

// func task(id int){
// 	fmt.Println("Doing task", id)
// }

// func main(){
// 	for i := 0; i <= 10; i++ {
// 	go func() { // adding go mean we are using gorutine
// 		fmt.Println(i)
// 	}()
// }
// time.Sleep(time.Second*2)
// }

func task(name string){
	for i:=1; i<=5; i++{
		fmt.Println(name, i)
		time.Sleep(300 * time.Microsecond)
	}
}

func main(){
	go task("A") // both function are run concurrently. task("A") and task("B")
	go task("B")
	time.Sleep(3 * time.Second)
}