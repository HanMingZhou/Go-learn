package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "10"
	var num int
	// string into int
	num, _ = strconv.Atoi(str)
	fmt.Printf("num,%T,%d\n", num, num)
	// int into string
	A := strconv.Itoa(num)
	fmt.Printf("A,%T,%s\n", A, A)

	x := 3.14
	strX := strconv.FormatFloat(x, 'f', 3, 32)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", x, strX)

	var i interface{} = "hello,world"
	str, ok := i.(string)
	if ok {
		fmt.Printf("i is %T", i)
	} else {
		fmt.Println("conversion failed")
	}

}
