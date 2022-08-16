// Package chapter1 章节1，类型与引用
package chapter1

import "fmt"

// Example2
// 值类型：int、float、bool、array、struct // 使用这些类型的变量直接指向存在内存中的值，所有非引用操作，传递的都是值本身
// 对应的指针类型：*int、*float、*bool 和 *string // 所有类型，都有其对应的指针类型，包括指针的指针类型
// 对应数组：[]int、[]float、[]bool 和 []string
// 对应指针数组：[]*int、[]*float、[]*bool 和 []*string
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

	// 见识一下指针类型
	//abarr := []bool{false, true}
	//var abarr1 = []*bool{}
	//println(abarr)
	//println(abarr1)
}

// 严格区别以下三种类型：
// 值类型，如 整型、浮点、布尔、array、struct 等
// 引用类型，如 string、slice、map、interface、channel、func 等
// 指针类型，值类型和引用类型，都可以获取其指针类型
