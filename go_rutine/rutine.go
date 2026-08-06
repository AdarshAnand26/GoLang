package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("Doing task", id)
}

func main(){
	for i := 0; i <= 10; i++ {
	go func() { // adding go mean we are using gorutine
		fmt.Println(i)
	}()
}
time.Sleep(time.Second*2)
}
