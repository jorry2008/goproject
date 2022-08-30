package chapter2

import (
	"fmt"
	"time"
)

// 管道是一种环形队列结构，用于协程间传递数据，在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。

// 这节使用一些demo，来说明 goroutine 和 channel 的关系
/*
一、在使用协程时，如果不基于管道阻塞特性，很可能导致新创建的协程无法执行完成而中断
二、设置缓冲的管道可以在同一个协程中同步操作，未设置缓冲的管道只能在不同协程之间进行通信，否则出现逻辑上的卡壳问题
三、chan 不允许并发读写，它基于 互斥锁 实现
四、chan 的 close 关闭问题？


*/

func Example21() {
	fmt.Println("Example21:")

	// 0.管道的本意：死锁问题
	// 管道用于协程间数据传递，有同步和阻塞两种模式
	// 0.1.同步模式（无缓存）在同一个协程下，当然无法正常工作
	//ch1 := make(chan int)
	//ch1 <- 5 // 这里发送了，但是程序就卡在这里了，因为必须等待接收方接收
	//_, ok := <-ch1 // 发送方因为没有接收方接收，所以程序永远无法执行到这一步【这个流程逻辑上就卡壳儿了】
	//println(ok) // 报错，所有协程都是死的。。。，这里使用了同步管道，确没有创建新的协程

	// 0.2.异步执行，可以在同一个协程中执行
	ch2 := make(chan int, 1) // 异步模式
	ch2 <- 50                // 因为有缓存，所以程序继续往下执行
	ch2v, ok2 := <-ch2
	println(ch2v, ok2)

	// 0.3.实现协程间数据传递
	ch3 := make(chan int) // 异步模式
	go func() {
		ch3 <- 51
	}() // 匿名函数，并立马调用
	ch3v, ok3 := <-ch3
	println(ch3v, ok3)

	// 1.新开协程被中断：当前主程序可能就已经将 say1() 执行完成并结束了主程序，而 say2() 由于所需时间较长，仍在执行而临时中断
	//go say2("say2 world")
	//say1("say1 hello")

	// 2.利用管道的阻塞特性：可保障新开的协程执行完毕
	ch22 := make(chan int, 1) // 以下是同步操作，这里必须设置缓存，否则就卡壳了
	say3("阻塞测试", ch22)        // 这里直接调用，没有创建新协程，所以仍然是同步操作
	say1("say1 hello")
	ch22v, ok := <-ch22
	println(ch22v, ok)

	// 3.演示管道的队列特性
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range + chan 天然就是一种阻塞的监听状态，通过 range 可以持续从 channel 中读取数据，类似于遍历；
	// 如果向此 channel 写数据的 goroutine 退出时，系统检测到这种情况后会 panic（系统中断），否则 range 将会永久阻塞；【理解一下 goroutine，可以认为是手动创建的 goroutine也可以当前执行的代码就是 goroutine】
	// 小结：range 对于 chan 而言，就是永久阻塞的，直到那个写入 channel 的那个协程 close() 管道遍历才会退出，否则 range 不会停止直到那个写入 channel 的那个协程挂了而抛出 panic 为止；
	for v := range c {
		fmt.Println("接收:", time.Now(), v)
	}

	// 4.单向通道
	ch := make(chan string) // 先创建一个双向通道
	go send(ch)
	receive(ch)
}
func say1(s string) {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("执行完成...")
}
func say2(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
}
func say3(s string, ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
	ch <- 0
	close(ch)
}
func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		fmt.Println("发送:", time.Now())
		c <- x
		time.Sleep(100)
		x, y = y, x+y
	}

	// 这里必须关闭
	close(c) // 关闭，表示告诉接收站，将以上数据接收完成后，就表示结束了，不需要在阻塞接收了！
}
func send(s chan<- string) { // 只能发送通道
	s <- "send发送的字条串：微客鸟窝"
}
func receive(r <-chan string) { // 只能接收通道
	str := <-r
	fmt.Println("receive接收到的字符串:", str)
}
