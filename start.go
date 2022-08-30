// 声明
// 风格要求：非注释的第一行必须指明当前文件所属的包（系统会自动格式化）
package main

import "github.com/jorry2008/goproject-demo/chapter3"

// 执行要求：package main 表示一个可独立执行的程序，每个 Go 应用程序都至少包含一个名为 main 的包，且必须包括执行入口 main 函数（这也是唯一一个强制包名不等于所在目录名一致的地方）
// 极简：Go 的 main 函数既没有参数，也没有返回值
func main() {

	// 第一单元

	// 1.声明与初始化
	//chapter1.Example1()

	// 2.类型与引用
	//chapter1.Example2()

	// 3.常量 枚举、私有常量、公有常量、iota
	//chapter1.Example3()
	//println(chapter1.Female)
	//println(chapter1.VAR_GLASS) // 严格区分大小写
	//println(chapter1.Var_glass)

	// 4.运算
	//chapter1.Example4()

	// 5.条件语句 if、switch、select
	//chapter1.Example5()

	// 6.循环语句 for、break、continue、遍历、goto
	//chapter1.Example6()

	// 7.函数、传值、传引用、回调函数、闭包、方法（函数和方法不同）
	//chapter1.Example7()
	//chapter1.Other()    // 调用其它包的公共函数，需要借用包来调用
	//chapter1.Fun1(1, 2) // 全局匿名函数，同样遵守首字母大写为公有函数

	// 8.作用域
	//chapter1.Example8()

	// 第二单元

	// 9.数组、多维、数组参数、浮点精度问题
	//chapter1.Example9()

	// 10.指针
	//chapter1.Example10()

	// 11.结构体
	//chapter2.Example11()
	//chapter2.Example11_2()

	// 12.切片
	//chapter2.Example12()

	// 13.集合
	//chapter2.Example13()

	// 14.遍历
	//chapter2.Example14()

	// 15.字符串和字节转化
	//chapter2.Example15()

	// 16.类型转换：自动转换，强制转换，精度问题
	//chapter2.Example16()

	// 17.递归写法
	//chapter2.Example17()

	// 18.接口
	//chapter2.Example18()

	// 19.错误处理及设计思路：error接口、panic()函数、recover()函数、defer语句、
	//chapter2.Example19()

	// 20.协程
	//chapter2.Example20()

	// 21.管道
	//chapter2.Example21()

	// 22.多路复用
	//chapter2.Example22()

	// 第三单元
	// 23.模块
	chapter3.Example23()
	// 24.包
	//chapter3.Example24()
	// 25.内建函数
	//chapter3.Example26()
	//chapter3.Example29()

	// 临时测试
	//chapter6.Example00007()

	// 第五单元：理解与总结

	// slice map struct 混合使用
	//chapter5.Example5_2()

	// var、make、new区别
	//chapter5.Example5_3()

	// 指针与引用
	//chapter5.Example5_4()

	//传值和传引用
	//chapter5.Example5_6()

	// 值接收器还是指针接收器
	//chapter5.Example5_7()

	// string 类型深入
	//chapter5.Example5_8()

	// 类型转换 强制、断言、显示、隐式
	//chapter5.Example5_10()

	// 接口的实现
	//chapter5.Example5_11()

	// 可寻址&不可寻址
	//chapter5.Example5_12()
	//chapter5.Example5_12_1()

}

// faq列表
// 1.什么叫隐式包，具体如何操作？
// 2.从类型的角度出发，整个Go语言表现出理解上的大统？鸭子类型
// 3.仍然不理解这个类型 var ppt uintptr？
