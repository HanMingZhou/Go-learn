package main

import "fmt"

func main() {

	var Iphone13 Iphone = Iphone{"A15"}
	var Nokia15 NokiaPhone = NokiaPhone{"885"}
	call(Iphone13)
	call(Nokia15)

}

// 接口
type Phone interface {
	// 定义call()方法
	call()
}

// iphone实现了接口
type Iphone struct {
	cpu string
}

func (iphone Iphone) call() {
	fmt.Println("i am iphone, cpu: ", iphone.cpu)
}

// Nokiaphone 实现了接口
type NokiaPhone struct {
	cpu string
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i am nokiaPhone, cpu: ", nokiaPhone.cpu)
}

func call(mobilephone Phone) {
	mobilephone.call()
}
