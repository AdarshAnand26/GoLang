package main

import (
	"fmt"
	//"time"
)

func main(){
	//Noraml switch type
	// i:=2
	// switch i{
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// default:
	// 	fmt.Println("default")
	// }

	//Multiple condition switch type
	// switch time.Now().Weekday(){ // Weekady and dayname should be start from capital letter
	// 	case time.Saturday, time.Sunday:
	// 		fmt.Println("Its weekend")
	// 	default:
	// 		fmt.Println("Its a weekday")
	// }

	//type switch
	whoiam:=func(i any){
		switch i.(type){
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Println("I am something else")	
		}	
	}
	whoiam(42)
	whoiam("hello")
	whoiam(3.14)
}
