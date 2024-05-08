package module1

import (
	"fmt"
	"go-learn/module2"
	"strings"
	// 模块module1要使用本地模块module2里的Add函数
	// 这里被import的本地模块的名称要和module2/go.mod里保持一致
)

func main() {
	a := 1
	b := 2
	sum := module2.Add(a, b)
	fmt.Printf("sum of %d and %d is %d\n", a, b, sum)

	fmt.Println("google" + " i can learn it")

	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	name := "John"
	age := 30
	height := 175.5

	result := fmt.Sprintf("Name: %s, Age: %d, Height: %.3f", name, age, height)
	fmt.Println(result)

	name = "rrrrrrrrrrr"
	fmt.Printf("%29v\n", name)

	a = 2

	for i := 0; i < 8; i++ {
		a = a * 2
		if a > 255 {
			fmt.Printf("i=%d\n", i)
			break

		}

	}
	fmt.Println(a)

	str := "这里是 www\n.runoob\n.com"
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.ToUpper(str)
	fmt.Println(str)

}
