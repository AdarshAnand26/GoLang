package main

import "fmt"

func main() {
	// var arr [5]int // giving 0 values on empty indexes
	// arr[0]=6
	// arr[1]=7
	// arr[4]=9
	// fmt.Println(arr)

	// var name [3]string  // give empty string on empty indexes
	// name[0]="suru"
	// name[1]="Adarsh"
	// fmt.Println(name)

	// var is [4]bool //give false when we not assign any value to it
	// fmt.Println(is)

	// //we can also declare and assign the value in arrays in a single line
	// var nums = [6]int{1, 2, 3, 4, 5, 6}
	// fmt.Println(nums)
	// var cars=[3]string{"BMW", "Audi", "Mercedes"}
	// fmt.Println(cars)

	//2D arrays
	var marks=[3][2]int{{},{6,7},{}}
	fmt.Println(marks)
}
