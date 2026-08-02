package main

 import "fmt"
//  import "slice"

 func main(){
	// var nums []int
	// fmt.Println(nums) // give a empty sice when did not assign any value to it

	// nums=append(nums,1)
	// nums=append(nums,2)
	// nums=append(nums,3)
	// fmt.Println(nums)

	// Another way to declare slice
	// var nums=make([]int,4,5) // here 4 is the length of slice and 5 is the capacity of slice
	// // nums[0]=1
	// // nums[1]=2
	// // nums[2]=3
	// // nums[3]=4
	// nums=append(nums,5) // its append after the last index of slice and increase the length of slice by 1
	// nums=append(nums, 6)
	// nums=append(nums, 7)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))// it will give the capacity of slice and it will automatically double when th elements added to it and its length is greater than its capacity

	// // Another way to declare slice
	// var nums=[]int {1,2,3,4,5}
	// fmt.Println(nums)

	// //copy function in slice
	// var nums=[]int{1,2,3,4,5}
	// var nums2=make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// //slice oprator
	// var nums=[]int{1,2,3,4,5}
	// fmt.Println(nums[1:4])
	// fmt.Println(nums[:4])
	// fmt.Println(nums[1:])

	// 	//slices
	// 	var nums1=[]int {1,2,3,4,5}
	// 	var nums2=[]int {1,2,3,4,5}
	// 	fmt.Println(slice.Equal(nums1, nums2)) // it will give an error because we can not compare slices in go
	
	//2D slices
	var nums=[][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(nums)

}
