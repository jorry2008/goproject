// 声明
package main

// 风格要求：非注释的第一行必须指明当前文件所属的包
// 执行要求：package main 表示一个可独立执行的程序，每个 Go 应用程序都至少包含一个名为 main 的包，且必须包括执行入口 main 函数

import (
	"demo/chapter1"
	"demo/chapter2"
	"demo/chapter6"
)

func main() {

	// 第一单元

	// 1.声明与初始化
	chapter1.Example1()

	// 2.类型与引用
	chapter1.Example2()

	// 3.常量 枚举、私有常量、公有常量、iota
	chapter1.Example3()
	println(chapter1.Female)
	println(chapter1.VAR_GLASS) // 严格区分大小写
	println(chapter1.Var_glass)

	// 4.运算
	chapter1.Example4()

	// 5.条件语句 if、switch、select
	chapter1.Example5()

	// 6.循环语句 for、break、continue、遍历、goto
	chapter1.Example6()

	// 7.函数、传值、传引用、回调函数、闭包、方法（函数和方法不同）
	chapter1.Example7()
	chapter1.Other()    // 调用其它包的公共函数，需要借用包来调用
	chapter1.Fun1(1, 2) // 全局匿名函数，同样遵守首字母大写为公有函数

	// 8.作用域
	chapter1.Example8()

	// 第二单元

	// 9.数组、多维、数组参数、浮点精度问题
	chapter2.Example9()

	// 10.指针
	chapter2.Example10()

	// 11.结构体
	chapter2.Example11()

	// 临时测试
	chapter6.Example00007()

}

// faq列表
// 1.什么叫隐式包，具体如何操作？
// 2.从类型的角度出发，整个Go语言表现出理解上的大统？
// 3.仍然不理解这个类型 var ppt uintptr？
