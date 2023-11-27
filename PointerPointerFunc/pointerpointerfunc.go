package main

import "fmt"

func main() {

	var a int = 10
	var b int = 33

	fmt.Println("before change, a、b=", a, b)
	Swap(&a, &b)
	fmt.Println("after changed, a、b=", a, b)

	var A int = 10000
	var B int = 33333
	A, B = B, A
	fmt.Println("after changed, A、B=", A, B)

}

func Swap(x, y *int) {
	*x, *y = *y, *x
}
