package chapter5

import "fmt"

/*
// go体系中，共有三大类型：值类型、引用类型、指针类型：
// 值类型，如 整型、浮点、布尔、array、struct 等
// 引用类型，如 string、slice、map、interface、channel、func 等
// 指针类型，值类型和引用类型，都可以获取其指针类型


new：
func new(Type) *Type
接收一个参数，这个参数是一种类型，而不是一个值，分配好内存后，返回一个指向该类型内存地址的指针，这个指针指向的内容的值为该类型的零值

对于不同的数据类型，零值的意义是完全不一样的
比如，对于bool类型，零值为false；int的零值为0；string的零值是空字符串
类型零值，就是初始化值，不同类型的值不同，结构体类型也是综合的

make：
func make(t Type, size ...IntegerType) Type
同样用于内存分配，但和 new 不同，make 用于 chan，slice 和 map 的分配，而且返回的类型就是这三个类型本身，而不是它们的指针，因为这三种类型本身就是引用类型，所以就没必要返回他们的指针了

具体而言，有如下几种用法：
（1）make(map[string]string)：即缺少长度的参数，只传类型，这种用法只能用在类型为 map 或 chan 的场景
（2）make([]int, 2)：指定了长度，例如make([]int, 2)返回的是一个长度为2的slice
（3）make([]int, 2, 4)：既指定了长度len为2，又指定了cap为4

注意：由于这三种类型都是引用类型，所以必须得初始化，但是并不是置为零值





make 用来创建 map、slice、channel
new 用来创建值类型

Golang 的引用类型包括 slice、map 和 chan，它们有复杂的内部结构，除了申请内存外，还需要初始化相关属性，make 会被编译器翻译 成具体的创建函数，由其分配内存和初始化成员结构，返回对象而非指针。
内置函数 new 计算类型大小，为其分配零值内存，返回指针。

new 和 make 均是用于分配内存：
new 用于值类型和用户定义的类型，如自定义结构体。
make 用于内置引用类型（切片、map 和管道）。
它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。
new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针。
make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作。（返回引用？）

零值？

参考：https://www.kancloud.cn/uvohp5na133/golang/933995

// 关于编译底层转换的问题？


还需要深入研究...

*/

func MakeNew() {
	// new 可以分配所有类型，并返回指针
	// 值类型、引用类型、指针类型，都可以使用 new 申请空间，并返回指针

	aaa := new(int)      // 值类型
	bbb := new(struct{}) // 值类型
	ccc := new([]int)    // 引用类型
	ddd := new(string)   // 特殊的引用类型
	eee := new(func())
	fff := new(interface{})
	ggg := new(chan string)
	hhh := new(*[]map[string]map[string]interface{}) // 无限添加。。。。。，返回指针的指针
	fmt.Println(aaa, bbb, ccc, ddd, eee, fff, ggg, hhh)

	// make 的使用范围
	// string、slice、map、interface、channel、func

	ss1 := make([]int, 10)
	ss2 := make(map[int]int)
	ss3 := make(chan string)
	//ss4 := make(func)
	//ss5 := make(interface)
	//ss6 := make(string)
	fmt.Println(ss1, ss2, ss3)
}
