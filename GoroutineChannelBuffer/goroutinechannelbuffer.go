package main

import (
	"fmt"
	"time"
)

func main() {

	channel_A := make(chan int, 5)
	go put(channel_A)

	for {
		time.Sleep(1 * time.Second)
		v, ok := <-channel_A
		if ok {
			fmt.Println("取出数据", v)
		} else {
			break
		}
	}

}

func put(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i
		time.Sleep(10 * time.Microsecond)
		fmt.Println("put channel success >-:", i)
	}
	fmt.Println("所有的都放进去了！关闭缓冲区，但是里面的数据不会丢失，还能取出。 ")
	close(c)
}
