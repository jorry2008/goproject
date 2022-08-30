package chapter3

import "fmt"

// 常用内建函数说明

func Example26() {
	fmt.Println("Example26:")

	// println() 和 fmt.Println() 两者的差异
	/*
		// 定义不同：
		println 属于 builtin 包，所以可以直接用，而 Println 属于 fmt 包，所以在没有用点来导入 fmt 的情况下（就是一般情况下），要 fmt.Println 这样调用。
		然后，看函数说明，
		println 是 func println(args ...Type)
		fmt.Println 是 func Println(a ...interface{}) (n int, err error)
		这个区别就很大了。
		两个函数其实都是可以接纳任何类型的对象，而且不限个数variadic，但是 println 没有返回值，而 fmt.Println 是有返回值的！
		Println 第一个返回值是 the number of bytes written，也就是往输出上写入了多少个字节，搞清楚一个输出字符串有多少字节很重要，特别是，比如，你想使用 bufio 这种包的时候，你会发现 unicode 什么的真的是恶心，对字节数心里没有数那真的是寸步难行，
		所以 fmt.Println 可以用来给你做实验搞清字节数规律什么的。。。而 err 含有可能的 error

		// 输出不同：
		fmt.Println 输出到标准输出，而 println 输出至标准错误
		println 主要程序启动和调试时用的，应该是语言内部实现主要用它

		// 两者的效果也不同，println 对结构体输出为指针（不接受数组和结构体），fmt.Println 输出的是类型的字面量且会调用各种类型的 String() string 或 Error() string 方法
		参考：https://www.zhihu.com/question/335186436/answer/756368792
	*/

	type multiMap map[string]map[string]string // 定义一个新类型（或叫声明一个新类型，如同声明一个新变量的意义一样）

	varmap := multiMap{"a": {"b": "b", "c": "c"}, "d": {"e": "e", "f": "f"}}
	println(varmap)
	fmt.Println(varmap)

	// 输出
	// 0xc0000724b0
	// map[a:map[b:b c:c] d:map[e:e f:f]]

	// 所以，在调试结构体时，建议使用 fmt.Println() 方法
}
