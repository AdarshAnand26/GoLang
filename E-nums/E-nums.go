package main

import "fmt"

// type orderStatus int 
// const(
// 	Recived orderStatus=iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )
// func changeOrderStatus(status orderStatus){
// 	fmt.Println("Change order status to", status)
// }

type orderStatus string 

const(
	Recived orderStatus="Recived"
	Confirmed orderStatus = "confirmed"
	Prepared orderStatus = "Prepared"
	Delivered orderStatus = "Delivered"
)

func changeOrderStatus(status orderStatus){
	fmt.Println("Change order status to", status)
}


func main(){
	changeOrderStatus(Recived)
}
