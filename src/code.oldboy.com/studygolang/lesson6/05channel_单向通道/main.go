package main

import "fmt"

// channel 单向通道
func main() {

	// 定义一个 容量为10 的存储类型为 int 的通道  (ch1 默认既能发送，又能接收值)
	var ch1 = make(chan int, 10)

	// 传入一个通道，指定它只能发送值 (单向通道)
	setValue(ch1)

	// 传入一个通道，指定它只能接收值 (单向通道)
	getValue(ch1)

}

// <-chan 代表只能从通道中接收值
func getValue(ch <-chan int) {
	v := <-ch // 从通道中接收值
	fmt.Println(v)

	// ch <- 10000  // 不能往通道里发送值： invalid operation: ch <- 10000 (send to receive-only type <-chan int)
}

// chan<-  代表是只能发送值到通道中
func setValue(ch chan<- int) {
	ch <- 2000 // 可以发送值到通道中。
	// v := <-ch   // 不能从通道中接收值 : invalid operation: cannot receive from send-only channel ch (variable of type chan<- int)
}
