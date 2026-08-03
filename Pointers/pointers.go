package main

import "fmt"

func changeNum(num int){
	num=5
	fmt.Println("In ChangedNum:", num)
	fmt.Println("Memory Address: ", &num)
}

func main(){
	num:=1
	changeNum(num)
	fmt.Println("After calling changedNum:", num)
	fmt.Println("Memory Address: ", &num)
}
//using above method its no changed the value of num


//now we use pointers to change the value of num
// func changeNum(num *int){
// 	*num=5
// 	fmt.Println("In ChangedNum:", *num)
// 	fmt.Println("Memory Address: ", &num)
// }

// func main(){
// 	num:=1
// 	changeNum(&num)
// 	fmt.Println("After calling changedNum:", num)
// 	fmt.Println("Memory Address: ", &num)
// }