package main

import "fmt"

func main() {

	A := bibao()
	re := A()
	fmt.Println(re)
}

func bibao() func() int {
	a := 1
	return func() int {
		a++
		return a
	}
}
