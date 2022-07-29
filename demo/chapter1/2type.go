// Package chapter1 章节1，类型与引用
package chapter1

import "fmt"

// Example2 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值，所有非引用操作，传递的都是值本身
func Example2() {
	println("Example2:")
	vara := 35
	// 取址运算符 &
	println(&vara) // 再判断一下是否是指针类型？
	// 解引用运算符 *
	println(*&vara)
	//println(reflect.TypeOf(&vara)) // (0x1572a0,0x120920)，暂时不懂.....

	// fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
	fmt.Printf("&vara 引用类型为 %T\n", &vara)
	fmt.Printf("*&vara 解引用类型为 %T\n", *&vara)

}

/*
如何理解 go 中的引用类型呢？
同 js 类型一样，只有 int float bool string 几个基本类型，在全局范围内正常传递的是值，
除此之外，其它类型都是传引用，有点类似对象的意思，
值是最抽象的一层，
引用是在值的基础之上的一种实现，
这也是 go 语言运行的基石。
*/
