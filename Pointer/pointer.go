package main

import "fmt"

func main() {
	a := 10

	p := &a

	fmt.Printf("p type is %T\n", p)

	fmt.Println("&a=", &a)

	fmt.Println("p=", p)
	fmt.Println("*p=", *p)
	fmt.Println("&p=", &p)

	var ptr *int

	fmt.Println("ptr", ptr)
	if ptr == nil {
		fmt.Println("Nil")
	} else {
		fmt.Println(ptr)
	}

	var arr [3]int = [3]int{1, 100, 200}
	var pter [3]*int

	for i := 0; i < len(arr); i++ {
		pter[i] = &arr[i]
	}
	for i := 0; i < len(pter); i++ {
		fmt.Printf("a[%d] = %d\n", i, *pter[i])
		fmt.Printf("a[%d] = %d\n", i, pter[i])
	}

	arrB := [3]int{10, 20, 30}
	ptrB := [3]*int{}        //指针数组
	var ptrC *[3]int = &arrB // 数组指针
	/*range循环中的x变量是临时变量range循环只是将值拷贝到x变量中。
	因此内存地址都是一样的。*/
	for i, x := range arrB {
		ptrB[i] = &x
	}
	for i, x := range ptrB {
		fmt.Println("ptrB i,x,*x", i, x, *x)
	}

	for i := 0; i < len(arrB); i++ {
		ptrB[i] = &arrB[i]
	}
	for i, x := range ptrB {
		fmt.Println("ptrB i,x,*x", i, x, *x)
	}

	//数组、指针数组 、数组指针
	for i := 0; i < len(arrB); i++ {
		fmt.Println(&arrB[i], ptrB[i], &ptrC[i])
	}
}
