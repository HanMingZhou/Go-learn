package main

import (
	"fmt"
)

func main() {
	AAA()
	BBB()
	CCC()

}

func AAA() {
	fmt.Println("FUNC AAA")
}

func BBB() {
	/*
				defer func() {
					if err := recover(); err != nil {
						fmt.Println("recover is func BBB")
						fmt.Println(fmt.Sprintf("%T %v", err, err))
					}

				}()


		defer func() { //不能捕获panic
			recoverDefault()

		}()
	*/
	defer recoverDefault() //defer recoverDefault()  //能捕获panic
	panic("FUNC BBB,能捕获panic")
}

func CCC() {
	fmt.Println("FUNC CCC")
}

func recoverDefault() {
	if err := recover(); err != nil {
		fmt.Println("recover is func BBB")
		fmt.Println(fmt.Sprintf("%T %v", err, err))
	}
}
