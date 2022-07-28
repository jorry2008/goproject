// Package chapter1 章节1，类型与引用
package chapter1

// Example2 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值，所有非引用操作，传递的都是值本身
func Example2() {
	println("Example2:")
	vara := 35
	println(&vara)
}
