package main

import "fmt"
import "maps"

func main(){

	// m:=make(map[string]int)
	// m["maths"]=90
	// m["science"]=80
	// m["english"]=70
	// delete(m,"science")
	// clear(m)
	// fmt.Println(len(m))

	//Another way to declare map
	// m:=map[string]int{
	// 	"maths":90,
	// 	"science":80,}

	// fmt.Println(m)

	//  _,ok:=m["science"] // check if key is present or not
	// if ok{
	// 	fmt.Println("Key is present")
	// }else{
	// 	fmt.Println("Key is not present")
	// }

	// v, ok:=m["bio"] // check if key is present or not
	// fmt.Println(v)
	// if ok{
	// 	fmt.Println("Key is present and value is",v)
	// }else{
	// 	fmt.Println("Key is not present")
	// }

	//map functions
	m1:=map[string]string{
		"name":"suru",
		"age":"20",
	}
	m2:=map[string]string{
		"name":"Adarsh",
		"age":"21",
	}	
	fmt.Println(maps.Equal(m1,m2)) // check if two maps are equal or not
	
}
