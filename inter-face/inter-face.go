package main

import "fmt"

type paymenter interface{
	pay(amount float32)
	refuned(amount float32, account_no string)
}

type payment struct{
	getway paymenter
}

func(p payment)makePayment(amount float32){
	p.getway.pay(amount)
}

type razorPay struct{}
func(r razorPay)pay(amount float32){
	fmt.Println("making payment using razorpay",amount)
}

type payPay struct{}
func(r payPay)pay(amount float32){
	fmt.Println("making payment using paypal",amount)
}

type stripe struct{}
func(r stripe)pay(amount float32){
	fmt.Println("making payment using stripe",amount)
}
func (s stripe) refuned(amount float32, account_no string) {
	fmt.Println("Refunded", amount, "to", account_no)
}

func main(){
	stripPaymentgw:=stripe{}
	newPayment:=payment{
		getway: stripPaymentgw,
	}
	newPayment.makePayment(100)
}
