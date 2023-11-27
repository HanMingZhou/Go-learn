package main

import "fmt"

func GetMax(num1, num2 int) int {

	//判断：num1 num2

	// 断言：value, ok := interface{}(a).(int)
	switch t := interface{}(num1).(type) {
	case int:
		fmt.Printf("the type of num1 is %T\n", t)
		goto JudgeNum2
	default:
		fmt.Println("type is string,please input right type(ints)")
	}

JudgeNum2:

	switch t := interface{}(num2).(type) {
	case int:
		fmt.Printf("the type of num2 is %T\n", t)

	default:
		fmt.Println("type is string,please input right type(int)")
	}

	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func main() {

	MAX := GetMax(1, 2)
	fmt.Printf("the MaxNumber between %v and %v is %v\n", 1, 2, MAX)

	/*
		11111
		333333
	*/
}
