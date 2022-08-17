// Package chapter5 这一节，包含：类型相关的理论 + 申请空间相关的实践，非常重要
package chapter5

import (
	"fmt"
)

// go体系中，共有三大类型：值类型、引用类型、指针类型：
// 1.值类型，如 整型、浮点、布尔、array、struct 等；
// 2.引用类型，如 string、slice、map、interface、channel、func 等；
// 3.指针类型，值类型和引用类型，都可以获取其指针类型；

// new()：
// func new(Type) *Type
// 接收所有类型，用于分配并初始化，返回一个指向该类型内存地址的指针，这个指针指向的内容的值为该类型的零值

// make()：
// func make(t Type, size ...IntegerType) Type
// 用于分配并初始化，make 只用于 chan，slice 和 map 的分配，而且返回的类型就是这三个类型本身（也叫对象？），而不是它们的指针，因为这三种类型本身就是引用类型，所以就没必要返回他们的指针了
// 注意：由于这三种类型都是引用类型，所以必须得初始化才可正常使用，否则只能接收同类型的其它值，使用 make 即完成了声明，也初始化为了零值

// 异：
// new 只分配内存并赋零值，对于所有类型都有效；
// make 只能用于 slice、map 和 channel 的分配内存及零值初始化；（它们三者，只声明类型，是不分配空间的且不可用，使用make没有初始值，但分配了空间，可用）

// 同：
// 两者都是内存的分配（堆上），对应的内存地址必须存在，就是必须有初始化值

// make() 的几种用法：
// （1）make(map[string]string)：即缺少长度的参数，只传类型，这种用法只能用在类型为 map 或 chan 的场景
// （2）make([]int, 2)：指定了长度，例如 make([]int, 2) 返回的是一个长度为 2 的 slice
// （3）make([]int, 2, 4)：既指定了长度 len 为 2，又指定了 cap 为 4

// 1.关于编译底层转换的问题？
// 2.在引用类型上使用 new，返回的是指针，如果*指针，得到的是引用？（引用是不能被解引用的，而且支持编译转换）？

func Example5_3() {
	fmt.Println("Example5_3:")

	// 注意：new() 对值类型和引用类型，操作的行为不一样，对值类型 new() 申请了空间可直接使用，对引用类型 new() 没有申请空间不能直接使用！

	// 一、new 可以分配所有类型，初始化零值，返回指针
	p1 := new(bool)     // 值类型，零值为 false
	p2 := new(int)      // 值类型，零值为 0
	p3 := new([4]int)   // 值类型，分配了4个空间的数组，且每个元素给零值 0【只声明才是空数组，这里不是空数组，有指针或引用，不可能是空数组的，因为需要指向】
	p4 := new(struct{}) // 值类型，有空间，零值为 {}，因为属性为空【实属意外，结构体是值类型 ^_^】
	fmt.Println(p1, p2, p3, p4)
	fmt.Println(*p1, *p2, *p3, *p4)
	// 对于 值类型的 指针使用，直接 *ptr 就可以了，比如：*p1=true  (*p3)[1]=25

	p5 := new(string) // 特殊的引用类型（需要专门讨论）

	p6 := new([]int)                                 // 引用类型，这里只是返回了一个指针，需要重新指向才可以使用
	p7 := new(map[string]string)                     // 引用类型
	p8 := new(chan string)                           // 引用类型
	p9 := new(func())                                // 引用类型，函数的入口就是函数名，它在内存中是个地址串，所以必须是引用类型
	p10 := new(interface{})                          // 引用类型，同上
	p11 := new(*[]map[string]map[string]interface{}) // 引用类型，指针当然是引用类型 // 这里可以无限添加，并返回指针的指针
	fmt.Println(p5, p6, p7, p8, p9, p10, p11)
	fmt.Println(*p5, *p6, *p7, *p8, *p9, *p10, *p11)
	// 引用类型的指针并不能直接使用，必须要初始化指向

	// make 的使用范围【make()怎么使用都不会错！！！！在 slice、map、channel 上御用】
	// string、slice、map、channel、interface、func

	ss1 := make([]int, 10)
	ss2 := make(map[int]int)
	ss3 := make(chan string)
	//ss4 := make(func)
	//ss5 := make(interface)
	//ss6 := make(string)
	fmt.Println(ss1, ss2, ss3)

	// var、new()、make() 测试
	VarNewMake()
}

// var、new、make
func VarNewMake() {
	// 一、var new() make() 操作切片 slice
	// 1.优先使用 make() 创建，并指定 len 或选用 cap，这种切片空间无限，可直接使用；
	// 2.可以使用 var 声明，但无法直接使用（能间接接收同类型）；nil
	// 3.可以使用 var 创建，必须要初始化给值，给多少值空间就是多大，不能超出；
	// 4.new() 只能创建对应类型的指针，没有值空间，无法直接使用（能间接接收同类型）；*解引用为 nil
	v1 := make([]int, 2, 4)
	v1[0] = 10 // 直接使用

	var v2 []int
	//v2[0] = 10 // 不能直接使用（没有空间）
	v2 = v1 // 间接接收同类型（这里也是算声明并初始化）
	v2[0] = 10

	v3 := []int{10} // 等于 var v3 = []int{10}
	v3[0] = 25      // 大小只有1，无法拓展

	v4 := new([]int)
	//(*v4)[0] = 10 // 不能直接使用（只有指针没有值空间）
	*v4 = v1
	v4 = &v1 // 间接接收同类型
	(*v4)[0] = 22

	// 二、var new() make() 操作集合 map
	// 1.优先使用 make() 创建，直接使用即可；
	// 2.可以使用 var 声明，但无法直接使用（能间接接收同类型）；nil
	// 3.可以使用 var 创建，必须初始化，即可使用；
	// 4.new() 只能创建对应类型的指针，没有值空间，无法直接使用（能间接接收同类型）；*解引用为 nil
	v5 := make(map[string]string)
	v5["a"] = "a" // 能直接使用

	var v6 map[string]string
	//v6["a"] = "a" // 不能直接使用（没有空间）
	v6 = v5 // 间接接收同类型（这里也是算声明并初始化）
	v6["a"] = "a"

	v7 := map[string]string{} // var v7 = map[string]string{}，给空值就行，map自身就是无限的
	v7["a"] = "a"

	v8 := new(map[string]string)
	//v8["a"] = "a" // 不能直接使用（只有指针没有值空间）
	*v8 = v5
	v8 = &v5 // 间接接收同类型
	(*v8)["a"] = "a"

	// 三、var new() make() 操作集合 chan
	// 1.使用 make() 创建，直接使用即可；
	// 2.new() 只能创建对应类型的指针，没有值空间，无法直接使用（能间接接收同类型）；*解引用为 nil
	v9 := make(chan int)
	//v9 <- 25 // 可直接用

	v10 := new(chan int)
	fmt.Println(*v10 == nil)
	*v10 = v9
	v10 = &v9 // 间接接收同类型
	//*v10 <- 25 // 可用
}

// 结论：
// 所有类型，都可以使用 new() 申请空间并初始化，进而实现类型的实例化，即通用又省空间（指针嘛），那为什么还要 var 和 make()？
// new() 一定会分配空间，且只能返回指针，而且对指针的操作是透明的，容易混淆出错，使用起来并不安全，实际使用中并不方便，注意：new()无法创建出直接可用的 slice、map、chan，它只创建了其指针类型，但并未给出空间和值；
// var 声明一个变量并指明变量的数据类型，可以声明（没有空间），并选择性初始化，var type 方式使用灵活、简单、自由，即可以创建值类型，又可以创建引用类型，但对于 slice、map、channel 而言；
// make() 在使用上是无可替代的！因为 make() 可以对 slice、map、channel 分配空间（强制的），并且 slice、map、channel 天然就是引用，返回引用类型本身，使用方便（能实现指针的效果，但更安全快捷）；
// make() 对 slice、map、channel 怎么用都不会错...！！！

// 如何选择？
// 大前提：所有类型，在方便的时候尽量使用 var 声明并可选初始化，必要时配合 new() 及 make() 操作；
// 对于 slice 必须使用 make() 创建，因为它要指定长度才可以使用；（var 实现的切片不可变）
// 对于 map 可以使用 var 声明并初始化，只声明的 map 无法使用，也可以使用 make() 操作；
// 对于 channel 只能使用 make() 操作（new()只能创建同步管道）；
// 对于 new() 能不用就不用；
// 最后：不管是 var只声明、new()无法指定长度 创建出来的类型，都可以间接使用，而 make() 怎么创建都不会错，它会强制给对应的类型设置参数；
