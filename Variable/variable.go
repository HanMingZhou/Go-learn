package main

import (
	"fmt"
	variableouter "go-learn/VariableOuter"
)

func main() {

	var n int
	fmt.Println("main.n= ", n)
	for n < 10 {
		fmt.Println("for.n = ", n)
		n++
	}
	fmt.Println("main.n=", n)

	var m int = 666
	fmt.Println("main.m= ", m)
	for m := 0; m < 10; m++ {
		fmt.Println("for.m = ", m)

	}
	fmt.Println("main.n=", m)

	x, y := 0, 100
	fmt.Println("variableouter.SumTest(x, y)=", variableouter.SumTest(x, y))
	fmt.Println("variableouter.OuterVariable= ", variableouter.OuterVariable)
}
