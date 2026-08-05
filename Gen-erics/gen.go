package main 

import "fmt"

// func printslice(items[] int){
// 	for _, item:=range items{
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items[] string){
// 	for _, item:=range items{
// 		fmt.Println(item)
// 	}
// }

// func main(){
// 	nums:=[]int{1,2,3}
// 	printslice(nums)
// 	S_data:=[]string{"Suru, Adarsh"}
// 	printStringSlice(S_data)
// } 
//Above code repeat the same code just with different datatype, we use generics to overcome it.

// func printslice[T any](items[] T){
// 	for _, item:=range items{
// 		fmt.Println(item)
// 	}
// }

// func printslices[T interface{}](items[] T){
// 	for _, item:=range items{
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice[T int| string | bool ](items[] T){
// 	for _, item:=range items{
// 		fmt.Println(item)
// 	}
// }

// func main(){
// 	nums:=[]string{"musu"}
// 	//printslice(nums)
// 	printslices(nums)
// 	S_data:=[]int{1,2,3,4}
// 	printStringSlice(S_data)
// } // this is way we can use generice in go , in which we can assign any datatype or all the datatype

//we can also use generic in struct
type stack[T any] struct{
	element []T
} 
func main(){
	myStack:=stack[string]{
		element: [] string{"Suru"},
	};
	fmt.Println(myStack)
}