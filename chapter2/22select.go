package chapter2

import (
	"fmt"
	"time"
)

// select 应用场景太多了，比如：可以实现多路复用，即同时监听多个 channel

func Example22() {
	fmt.Println("Example22:")

	// 1.超时方案
	// select 巧妙的解决 goroutine 基于管道共享数据时的超时问题
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2) // 2秒之后将数据写入管道
		c1 <- "result 1"
	}()
	// 两个 case 以阻塞的模式轮询
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1): // 1秒后，定时器向当前协程释放一个时间
		fmt.Println("timeout 1")
		//case t := <-time.After(time.Second * 1): // 1秒后，定时器向当前协程返回事件
		//	fmt.Println("timeout 1", t)
	}

	// 2.多路复用（永久执行）
	// select 可以实现多路复用，即同时监听多个 channel
	// 发现哪个 channel 有数据产生，就执行相应的 case 分支
	// 如果同时有多个 case 分支可以执行，则会随机选择一个
	// 如果一个 case 分支都不可执行，则 select 会一直等待
	/*
		ch := make(chan int, 1) // 同协程，这里必须要缓冲
		for {
			select {
			case <-ch: // <- ch：接收通道，但是对值不处理
				time.Sleep(time.Second)
				fmt.Println("case 1 invoke")
			case data := <-ch: // data := <-ch：接收通道，并处理从通道中得到的结果
				time.Sleep(time.Second)
				fmt.Printf("case 2 invoke %d\n", data)
			case ch <- 100: // ch <- 100：向通道中发送数据【这个case肯定是无条件执行的】
				time.Sleep(time.Second)
				fmt.Println("case3 invoke")
			}
		}
	*/
	// 上面的程序运行起来之后，case 3 会首先执行，然后 case1 和 case2 会随机执行一个，程序就这样一直交替运行下去

	// 3.双向管道（永远执行，手动中断）
	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("显示：", <-c2) // 同步处于阻塞状态
		}
		quit <- 0 // 反向给发送方中断信号
	}()
	x, y := 0, 1
	// 处于长期监听状态
	for { // for + select 一起使用，非常像 range 所实现的功能
		select {
		case c2 <- x: // 注意：case 可以是 send 语句，也可以是 receive 语句，亦或者 default【管道发送操作】
			x, y = y, x+y
		case <-quit: // 这里根本不在乎返回了什么内容，只要 quit 管道中有信号，就表示中断【管道接收操作】
			fmt.Println("收到中断信号，quit...")
			return // 中断 for 循环
		}
	}
	// select 的 case 只关注对管道的写和读操作，即操作本身就是其触发事件

	// 4.解除死锁（永远执行）
	ch4 := make(chan int)
	for {
		select {
		case <-ch4:
			fmt.Println("case invoke")
		default:
			time.Sleep(time.Second)
			fmt.Println("default invoke") // 如果没有 default，这段程序只有一个 case，而且这个case永远也接收不到管道数据，因此死锁
		}
	}
}
