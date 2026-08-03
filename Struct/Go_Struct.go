package main

import (
	"fmt"
	"time"
)

type customer struct{
	name string
	phone string
}

type order struct{
	id string
	amount float32
	status string
	createdAt time.Time
	customer //struct embedding 
}

// declaring contructors 
func newOrder(id string, amount float32, status string)*order{
	myorder:=order{
		id:id,
		amount:amount,
		status:status,
	}
	return &myorder
}

// declaring methods
func(o*order)changeStatus(status string){
		o.status=status
	}
func(o order)getAmount()float32{
		return o.amount
	} 



func main(){

 myorder:=order{
	id:"1",
	amount:50.00,
	status:"reciced",
	createdAt: time.Now(),
 }

 myorder2:=order{
	id:"2",
	amount:40.00,
	status:"reciced",
	createdAt: time.Now(),
 }

 myorder4:=order{
	id:"4",
	amount:70.00,
	status:"reciced",
	createdAt: time.Now(),
	customer: customer{ //using struct embedding 
		name: "Adarsh",
		phone: "3846928164",
	},
 }

 //Also can we use like time
 newCustomer:=customer{
	name:"suru",
	phone: "2394892472",
 }

 myorder5:=order{
	id:"5",
	amount:80.00,
	status:"reciced",
	createdAt: time.Now(),
	customer: newCustomer,
 }



 //declaring sturct here only but can be use only once
 language:=struct{
	name string
	isGood bool
 }{"goLang",true}

 fmt.Println(language)

 myorder.changeStatus("confirmed")

 myorder3:=newOrder("3", 30.00, "recived")

 fmt.Println(myorder)
 fmt.Println(myorder2)
 fmt.Println(myorder3)
 fmt.Println(myorder4)
 fmt.Println(myorder5)


 fmt.Println("Amount:", myorder.getAmount())
}
