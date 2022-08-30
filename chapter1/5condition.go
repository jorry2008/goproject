package chapter1

import "fmt"

// Example5 注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断，更不支持合并运算符 ??
// 重点：Go 中只有 if、switch 和 select 三种条件语句（其中 select 只用于信道场景）
func Example5() {
	fmt.Println("Example5:")

	// 判断条件为 expr 逻辑运算表达式，标准的 c 语言结构，同 php
	if true {
		println("布尔表达式")
	} else if 1 == 1 && 8 == 63 {

	} else {

	}

	switch 5 {
	case 6:
		println("条件语句 if 和 switch 和 C 语言 php 语言一样，包括格式，只是多了一个 fallthrough，不支持三目运算符")
		break
	case 7:
		//statement(s)
		break
		// switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case（所以break在switch中可以省略），如果我们需要执行后面的 case，可以使用 fallthrough
	case 60:
		//statement(s)
		fallthrough // fallthrough 后的 case 不论是否匹配都会强制执行！！
	case 88:
		//statement(s)
		break
	default:
		//
	}

	// 专用于协程 -> 信道/通道
	// 遍历所有 case 信道，直接获取的数据并执行 case 为止
	// 在运行 select 时，会遍历所有（如果有机会的话）的 case 表达式，只要有一个信道有接收到数据，那么 select case 就执行并结束
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c2 <- "hello" // 通信，向信道中发送一条数据

	select {
	case msg1 := <-c1: // 表达式
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	default:
		fmt.Println("No data received.")
	}

}
