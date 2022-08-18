package chapter5

import (
	"fmt"
)

// 很多人无法清晰的知道 传值、传引用？
// 这个问题主要是因为语言不同，形参和实参的传递行为也不同，通常所说的传引用是针对 c、c++、java 而言的。
// 先说结论：在 Go 语言中，没有传引用这个东西。

// 值类型：int、float、bool、array、sturct
// 引用类型：slice、map、channel、interface、func、（特殊）string
// 指针类型：为了好理解将它单独出来（在归类中，它属于引用体系，叫透明引用）

// go只有传值与php一样，别多想，就是这么简单
// 注意：在测试值或引用传递的过程中，小心一点编译自动类型转换

type Cat struct {
	Name    string                     // 默认值 ""
	Age     int                        // 默认值 0
	Color   string                     // 默认值 ""
	Teacher *string                    // 默认值 nil
	cb      func(map2 *map[string]int) // 默认值 nil // 回调函数
}

func Example5_6() {
	fmt.Println("Example5_6:")

	// 自定义3个数据类型（只声明未初始化）
	var varInt int
	var varSlice []int
	var varCat Cat // 自定义类型同样有零值（即每个属性自身的默认值）

	SetInt(varInt) // 传递过去的就是值类型
	fmt.Println(varInt)
	SetIntPoint(&varInt) // 传递过去的就是值类型指针
	fmt.Println(varInt)

	varSlice = make([]int, 8, 16)
	SetSlice(varSlice) // 传递过去的是引用
	fmt.Println(varSlice)
	SetSlicePoint(&varSlice) // 传递过去的是引用指针
	fmt.Println(varSlice)

	SetCat(varCat) // 传递过去的是值类型
	fmt.Println(varCat)
	SetCatPoint(&varCat) // 传递过去的是值类型指针
	fmt.Println(varCat)

	// 小节：自定义结构体类型的用法与普通类型，完全一样！
}

func SetInt(i int) {
	i = 88
}

func SetIntPoint(i *int) {
	*i = 89
}

func SetSlice(s []int) {
	s[0] = 9
}

func SetSlicePoint(s *[]int) {
	(*s)[0] = 99
}

func SetCat(c Cat) {
	c.Name = "Tom"
	// (&c).Name = "ee" // 这个用法非常有意思，c是结构体，有属性，&c 为指针值类型，(&c).Name这个调用就是将值类型转化为指针值类型，然后再自动类型转换为值类型....
}

func SetCatPoint(c *Cat) {
	// (*c).Name = "Tom" // 根据类型，这个用法才是正常的
	c.Name = "Tom" // 可以直接访问属性，是基于自动类型转换
	c.Age = 50
	c.Color = "red"
}

// 经过测试，函数的形参是什么类型，实参对应就是什么类型，而且实参接收到的就是形参内存里储存的内容！！！！就是值传递
