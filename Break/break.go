package main

import "fmt"

func main() {

label:
	for i := 1; i < 10; i++ {
		for i2 := 4; i2 <= i; i2++ {
			if i2 > 4 {
				fmt.Println("when break is happening,the i =", i)
				break label

			}
		}

	}

	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
		fmt.Println("case Monday")
		break
	case "Tuesday":
		fmt.Println("It's Tuesday.")
		break // 跳出 switch 语句
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	}
}
