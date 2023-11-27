package main

import (
	"fmt"
)

func Chann(Ch chan int, StopChan chan bool) {
	for J := 0; J < 10; J++ {
		Ch <- J
	}
	StopChan <- true
}

func main() {

	ch := make(chan int)
	stopChan := make(chan bool)
	c := 0

	go Chann(ch, stopChan)

	for {
		select {
		case c = <-ch:
			fmt.Println("receive c", c)
		case s := <-ch:
			fmt.Println("receive s", s)
		case _ = <-stopChan:
			fmt.Println("两个通道都没有可用的数据,执行 default 子句中的语句")
			goto end
		}
	}
end:
	/*

		channelA := make(chan string)
		channelB := make(chan string)

		go func() {
			channelA <- "from channelA"
		}()

		go func() {
			channelB <- "from channelB"
		}()

		for { // 使用 select 语句非阻塞地从两个通道中获取数据
			time.Sleep(2 * time.Second)
			select {
			case messageA := <-channelA:
				fmt.Println(messageA)

			case messageB := <-channelB:
				fmt.Println(messageB)
			default: // 如果两个通道都没有可用的数据，则执行这里的语句
				fmt.Println("ther is no message from channelA or channelB")
			}

		}
	*/

}
