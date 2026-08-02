package main

import "fmt"

func main(){
	//nums:=[]int{1,2,3,4,5,6,7,8,9}
	// for i, v := range nums {
	// 	fmt.Printf("Index: %d, Value: %d\n", i, v)
	// }

	// sum:=0
	// for num:=range nums{
	// 	sum=sum+num
	// }
	// fmt.Println(sum)

	//Using map
	m:=map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	//Using string
	// for i,c:=range "compatation"{
	// 	fmt.Println(i,c)
	// }

	for i,c:=range "compatation"{
		fmt.Println(i,string(c))
	}

}
