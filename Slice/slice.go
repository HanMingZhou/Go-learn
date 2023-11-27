package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 注意cap容量的变化 */
	numbers = append(numbers, 2)
	printSlice(numbers)
	numbers = append(numbers, 3)
	printSlice(numbers)
	numbers = append(numbers, 4) // 可以看出，容量不够时，cap会自动扩容到2倍
	printSlice(numbers)

	numbers = append(numbers, 4, 5, 6, 7) // 可以看出，容量不够时，cap会自动扩容到2倍
	printSlice(numbers)

	a := []int{1, 2, 3, 4}
	printSlice(a)

	c := a
	printSlice(c)

	a = append(a, 5, 6)
	printSlice(a)
	printSlice(c) //append 会让切片和与他相关的切片脱钩

	d := &a
	printSlice(*d)

	a = append(a, 7, 8, 9)
	printSlice(a)
	printSlice(*d)

	x := make([]int, 4)
	b := x[:2]
	b[0] = 3
	fmt.Println("b", b)
	fmt.Println("x", x)

	y := make([]int, 4)
	k := y
	c[0] = 1
	fmt.Println(k)
	fmt.Println(y)
	k = append(k, 2)
	k[0] = 0
	fmt.Println(k)
	fmt.Println(y)
	fmt.Printf("切片b的地址：%p\n", k)
	fmt.Printf("切片y的地址：%p\n", y)
	/*脱钩的情况是由于切片底层数组扩张
	（创建了新数组替换旧数组）导致*/
	fmt.Println()

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
