package main

import "fmt"

func recursion(n int) int {
	if n > 0 {
		fmt.Println("n=", n)
		return n * recursion(n-1)
	}

	return 1
}

const LIM = 10

func main() {

	// 阶乘
	var n int = 2
	fmt.Printf("%v 阶乘 = %v\n", n, recursion(n))

	// Fibonacci
	//result := 0
	//var array []int
	var array [LIM]int
	for i := 0; i < LIM; i++ {
		array[i] = fibonacci(i)
		//result = fibonacci(i)
		//array = append(array, result)
		//fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	fmt.Println(array)

}

/*指的是这样一个数列：1、1、2、3、5、8、13、21、34
⋯⋯这个数列从第3项开始，每一项都等于前两项之和。*/
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
