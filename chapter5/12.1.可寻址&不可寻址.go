package chapter5

import (
	"fmt"
	"unsafe"
)

type person struct {
	name string
}

func getP() person {
	return person{name: "diffcoder"}
}

func getStr() string {
	return "diffcoder"
}

func getInt() int {
	return 1024
}

func Example5_12_1() {
	fmt.Println("Example5_12_1:")

	// 什么叫可寻址?
	// 可直接使用 & 操作符取地址的对象，就是可寻址的，这是一个检测手段（不要与实例化一个结构体的指针的操作混淆了）

	// 一、哪些是可以寻址的
	// 1.变量：&x
	name := "diffcoder"
	fmt.Println(&name) // 0xc00003c1f0

	// 2.指针：&*x
	fmt.Println(unsafe.Pointer(&person{"diffcoder"})) // 0xc00003c1f0 【unsafe.Pointer()取结构体的入口地址值】

	// 3.数组元素索引: &a[0]
	s1 := [3]int{1, 2, 3}
	fmt.Printf("%p\n", &s1[0]) // 0xc00000e3c0 // %p 显示地址值

	// 4.切片：[]type{}[]
	fmt.Println([]int{1, 2, 3}[1:]) // [2 3]，注意这个用法的原由

	// 5.切片元素索引：&s[0]
	s2 := []int{1, 2, 3}
	fmt.Println(&s2[0]) // 0xc00000e3c0

	// 二、哪些是不可以寻址
	// 1.常量
	const PI = 3.14
	//fmt.Println(&PI) // cannot take the address of PI

	// 2. 字符串
	//fmt.Println(&getStr())        // cannot take the address of getStr()
	//fmt.Printf("%p\n", &getStr()) // cannot take the address of getStr() // %p 取指针地址值

	// 3.函数或方法
	//fmt.Println(&getStr) // cannot take the address of getStr
	//fmt.Printf("%p\n",&getStr) // cannot take the address of getStr

	// 4.基本类型字面量
	// 原因：字面量分 基本类型字面量 和 复合型字面量， 基本类型字面量，是一个值的文本表示，都是不应该也是不可以被寻址的
	//fmt.Println(&getInt()) // cannot take the address of getInt()

	// 5.组合字面量: &struct{X type}{value}
	// 所有的组合字面量都是不可寻址
	// 所谓的组合字面量其实就是把对象的定义和初始化放在一步完成（创建并初始化），包括 结构体、数组、切片 和 map 各自的常规方式和组合字面量方式
	//fmt.Println(&getP()) // cannot take the address of getP()
	// 注意上面写法与这个写法的区别，下面这个写法代表不同意思，其中的 & 并不是取地址的操作，而代表实例化一个结构体的指针
	fmt.Println(&person{name: "diffcoder"}) // &{diffcoder}，这个其实是指针，可寻址的，注意：这种形式需要使用函数间接测试，否则会与实例化指针混淆
	// 虽然组合字面量是不可寻址的，但却可以对组合字面量的字段属性进行寻址（直接访问）这里比较特殊，请注意
	fmt.Println(getP().name) // diffcoder

	// 6.数组字面量
	// 数组字面量是不可寻址的，当你对数组字面量进行切片操作，其实就是寻找内部元素的地址，下面这段代码是会报错的
	//fmt.Println([3]int{1, 2, 3}[1:]) // invalid operation [3]int literal[1:] (slice of unaddressable value)
	//fmt.Printf("%p\n", &[3]int{1, 2, 3}[0]) // cannot take the address of [3]int literal[0]
	// 其实基本上要达到寻址，就通过使用一个变量承接一下，比如：
	arr := [3]int{1, 2, 3}
	fmt.Println(arr[1:])        // [2,3]
	fmt.Printf("%p\n", &arr[0]) // 0xc00000e3c0

	// 7.所有字面量都不可寻址
	// 在 Go 中内置的基本类型有：
	// 布尔类型：bool
	// 语言中的 11 个内置的整数数字类型：int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint 和 uintptr
	// 浮点数类型：float32 和 float64
	// 复数类型：complex64 和 complex128
	// 字符串类型：string
	// 而这些基本类型值的文本，就是基本类型字面量
	// !!!字面量，说白了就是未命名的常量，本质跟常量一样，他是不可寻址的
	//const VAR_STR = "var_str1"
	//"var_str2"
	// 以上两者，是一样的，它们只是没有被变量索引的文本而已，没有啥特性

	// 8.map 中的元素
	// 原因：字典比较特殊，可以从两个角度来反向推导，假设字典的元素是可寻址的，会出现什么问题？ 如果字典的元素不存在，则返回零值，而零值是不可变对象，如果能寻址问题就大了，而如果字典的元素存在，考虑到 Go 中 map 实现中元素的地址是变化的！！！这意味着寻址的结果也是无意义的，基于这两点，Map 中的元素不可寻址，符合常理
	//m1 := map[string]int{"diffcoder": 20}
	//fmt.Println(&m1["diffcoder"])        // cannot take the address of m["diffcoder"]
	//fmt.Printf("%p\n", &m1["diffcoder"]) // cannot take the address of m["diffcoder"]
	// 一个经典实例：
	//m2 := map[int]person{
	//	1: {name: "diffcoder"},
	//	2: {name: "tom"},
	//}
	//m2[1].name = "James" // cannot assign to struct field m[1].name in map

	// 如果需要改结构体里面的属性，可以改为如下：
	m3 := map[int]*person{ // 结构体是指针
		1: &person{name: "diffcoder"},
		2: &person{name: "tom"},
	}
	fmt.Println(m3[1].name) // diffcoder
	m3[1].name = "James"
	fmt.Println(m3[1].name) // James
}
