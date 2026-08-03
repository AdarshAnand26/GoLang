package main

import "fmt"

func add(nums...int)int{
	total:=0
	for _ , num:=range nums {
		total=total+num
	}
	return total

}

func main(){
	result:=add(1,2,3,3,4,4,55,6)
	fmt.Println(result)
}


