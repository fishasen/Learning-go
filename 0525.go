package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
	Keyboard string
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Printf("I am Nokia,I can call you!%s", nokiaPhone.Keyboard)
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am IPhone,I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	//phone.Keyboard = "s"
	phone.call()

	phone = new(IPhone)
	phone.call()
}
