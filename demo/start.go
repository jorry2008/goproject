// 声明
package main

// 风格要求：非注释的第一行必须指明当前文件所属的包
// 执行要求：package main 表示一个可独立执行的程序，每个 Go 应用程序都至少包含一个名为 main 的包，且必须包括执行入口 main 函数

import (
	"demo/chapter1"
)

func main() {

	// 1.声明与初始化
	chapter1.Example1()

	// 2.类型与引用
	chapter1.Example2()

	// 3.常量
	chapter1.Example3()
	println(chapter1.Female)
	println(chapter1.VAR_GLASS) // 严格区分大小写
	println(chapter1.Var_glass)

	// 函数
	chapter1.Example8888()
}
