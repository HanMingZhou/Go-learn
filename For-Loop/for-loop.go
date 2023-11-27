package main

import (
	findprime "FindPrime"
	"fmt"
)

func main() {

	map1 := make(map[int]float32)
	map1[0] = 1
	map1[1] = 1111
	map1[2] = 11111
	map1[10] = 11111

	strings := []string{"菜鸟教程", "google"}
	for index, val := range strings {
		fmt.Println(index, strings[index])
		fmt.Println(val)
	}

	numbers := []int{111, 2343, 5343, 53413, 112544}
	for i := range numbers {
		fmt.Println(numbers[i])
	}

	for i := range map1 {
		fmt.Println(i, map1[i])

	}

	/*
		//练习：编写一个函数，输出100以内的所有素数。每行只显示5个，并求和，
		//判断：素数为大于1且只能被自己整除
		var FLAG bool = false //初始值默认为素数
		Y := 2
		for Y <= 100 {
			//重新赋值，从2 ～ 100分别取出i，与从2 ～ (i-1) 的数中取出的j 进行比较，
			//是否被取出的数可以整除
			FLAG = true
			for j := 2; j < Y; j++ {
				if Y%j == 0 {
					FLAG = false
				}
			}
			if FLAG {
				fmt.Printf("%v \t", Y)

			}

			Y++
		}
	*/

	// 9X9乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%v X %v = %v\t", i, j, i*j)
		}
		fmt.Println()
	}

	var A int = 2
	var list []int
	flag := true // is prime defaultly
	count := 0

	for ; A <= 100; A++ { // 外层循环...
		flag = true
		/*if A is not prime, the flag is false;
		so the outside circul is always false,
		so the flag must be true*/
		//一个数的因子不可能超过它的一半加一

		for B := 2; B < A; B++ {
			if A%B == 0 {
				flag = false

				break
			}
		}

		if flag {
			list = append(list, A)
			count++
		}

	}

	fmt.Println(list)
	fmt.Println("the number of prime is ", count)

	// 查找X以内的素数
	fmt.Println("input the number you want to find its prime")
	var X int
	fmt.Scanln(&X)
	_, Primecount := findprime.FindPrime(X)
	fmt.Printf("the prime'number of %v is %v\n", X, Primecount)

	// 查找X以内的素数 with Goto
	fmt.Println("------PrimeWithGoto-----")
	fmt.Println("input the number you want to find its prime")
	var Y int
	fmt.Scanln(&Y)

	PrimeWithGotoList, PrimeWithGotoCount := findprime.FindPrimeWithGoto(Y)

	fmt.Printf("the prime'number of %v is %v\n", Y, PrimeWithGotoList)
	fmt.Printf("the %v has %v\n", Y, PrimeWithGotoCount)

}
