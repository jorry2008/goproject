package chapter1

// 局部声明或者初始中 := 操作很常用，也很方便
func Example3() {
	// 空白标识符，是一个只写变量，你不能得到它的值，正因为这种特性，空白标识符可用于接受多余的值并执行了初始化，但又可以不在当前作用域下使用，从而避免了错误（局部声明变量不使用报错的问题！！）
	_, numb, strs := numbers() //只获取函数返回值的后两个
	println(numb, strs)
}

// 一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
