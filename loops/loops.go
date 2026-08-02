package main 

import "fmt"

func main(){

	//while loop
	// i:=0
	// for i<=10{
	// 	fmt.Println(i)
	// 	i++
	// }

	// for loop
	// for j:=0; j<10; j++{
	// 	fmt.Println(j)
	// }

	//range
	for i:=range 9{
		if i==5{
			continue
		}
		fmt.Println(i)
		
	}
}