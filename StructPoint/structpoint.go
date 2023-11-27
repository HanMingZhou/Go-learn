package main

import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

type book struct {
	title  string
	author string
	id     int
}

func main() {
	var c Circle
	fmt.Println(c.radius)
	c.radius = 10.00
	fmt.Println(c.getArea())
	c.changeRadius(20)
	fmt.Println(c.radius)
	change(&c, 30)
	fmt.Println(c.radius)

	c2 := getAera2(c)
	fmt.Println("c2.redius ", c2)
	c3 := getArea3()
	fmt.Println("c3.redius ", c3)

	var book1 book = book{"Google", "Ym", 0}
	fmt.Println("book1", book1)
	var b *book = &book1
	b.title = "Micsoft"
	b.author = "My"
	b.id = 1
	fmt.Println("b", b)
	fmt.Println("*b", *b)
	fmt.Println("book1", book1)

}

func (c Circle) getArea() float64 {
	return c.radius * c.radius
}

// 注意如果想要更改成功c的值，这里需要传指针
func (c *Circle) changeRadius(radius float64) {
	fmt.Printf("c.changeRadius(20) is be executized  ")
	c.radius = radius
}

// 以下操作将不生效
// func (c Circle) changeRadius(radius float64) {
// 	fmt.Printf("c.changeRadius(20) is not be executized  ")
// 	c.radius = radius
// }

// 引用类型要想改变值需要传指针
func change(c *Circle, radius float64) {
	c.radius = radius
}

func getAera2(c Circle) Circle {
	var temp Circle
	temp.radius = c.radius * 10
	return temp
}

func getArea3() Circle {
	var temp Circle
	temp.radius = 0.9999
	return temp
}
