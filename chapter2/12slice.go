package chapter2

import (
	"fmt"
)

// 切片不是数组，数组一切是切片，两者都基于索引， 且键有排序（切片是指向底层的数组，切片是数组的抽象）
// 创建切片一般用 make 方法，参数1是指向的切片类型，参数2是存放元素的 len 个数，参数3是存放 cap 容量（个数的边界）
// 在切片中，存放的 len 个数超过 cap 容量，系统会从分配内存地址 (容量在原来的基础上 * 2)，len 方法获取长度，cap 获取容量

var emptySlice []int            // 声明一个切片，默认为 nil
var emptySlice1 = []int{}       // 声明一个切片，并初始化
var emptySlice2 []int = []int{} // 声明一个切片，同上

func Example12() {
	println("Example12:")

	// 只声明不初始化的切片默认为nil，表示还未分配空间，声明并使用{}初始化为空切片其值就是空
	println(emptySlice)                  // 空表示：[0/0]0x0
	println("返回内容:", emptySlice == nil)  // true【没有分配存储空间】
	println(emptySlice1)                 // [0/0]0xf553a0
	println("返回内容:", emptySlice1 == nil) // false【分配了存储空间，但没有值】

	s := []int{}                      // 切片
	s = append(s, 52, 50, 40, 30, 20) // 给切片添加元素
	//s[len(s)] = 45                              // 索引的方式增加一个元素
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // 数组

	fmt.Printf("%T, %T \n", s, a) // []int [10]int，从类型可以看出，切片没有长度问题，而数组的长是固定的

	// 数组使用角标范围读取（基于数组生成切片）
	c := a[3:5] // 从角标为 star 到 角标为 end-1 的元素（不包括 end 角标）
	for ci, civ := range c {
		println("c", ci, civ)
	}
	d := a[3:] // 从索引 3 起，到最后所有
	for di, div := range d {
		println("d", di, div)
	}
	e := a[:5] // 从索引 0 起，不包括索引 5
	for ei, eiv := range e {
		println("e", ei, eiv)
	}
	fmt.Printf("通过角标截取的新数组：%T \n", e)  // []int，使用角标截取数组，返回的是切片!其容量就是原数组的数组大小
	println("e大小和容量：", len(e), cap(e)) // cap(e) 就等于 len(a)

	// 切片的 make 操作，切片使用角标范围读取
	ms1 := make([]int, 5)     // len 必填，cap 选填
	ms2 := make([]int, 5, 12) //参数1表示存储类型，表示2存储数组长度，参数3是指切片最大长度（也叫容量），如果长度超出容量，他就会翻倍，分配一个长度 容量*2 的内存块，容量参数不设置，最大容量就是数组长度

	println("ms大小和容量：", len(ms1), cap(ms1)) // 没有指定容量时，容量和长度保持一致（动态的，也就是没有容量这个限制了）
	for msi, msv := range ms1 {
		println("ms：", msi, msv)
	}

	// 如果设置了 cap 容量，当超出长度时，切片的存储空间会按照cap容量来翻倍
	println("ms2大小和容量：", len(ms2), cap(ms2)) // len为5，cap为12
	ms2 = append(ms2, 10, 21, 25, 24, 23, 22, 33, 12, 14, 74, 44)
	println("拓充后的ms2大小和容量：", len(ms2), cap(ms2)) // len为16，cap为24
	for ms2i, ms2v := range ms2 {
		println("ms2", ms2i, ms2v)
	}

	// 角标截取，会返回新的切片，且索引会重置，容量继承自被截取的切片
	msm := ms2[10:13]
	println("msm长度和容量", len(msm), cap(msm)) // msm 的容量取 ms2 的最大角标加1即：end + 1
	for msmk, msmv := range msm {
		println("msm", msmk, msmv)
	}
	fmt.Printf("%T \n", msm) // []int

	// 数组和切片的初始化
	sss := []string{1: "ss", 3: "sss"}
	for sssk, sssv := range sss {
		println("sss", sssk, sssv)
	}
	//错误用法 sss[4] = "字符串"                          // 重点，切片支持自动拓展容量，但不能以这种方式，而是要使用 append 这种函数
	println("sss长度和容量", len(sss), cap(sss)) // 支持初始化指定索引赋值，整个切片的长度和容量，取决于最大的角标+1

	aaa := [5]string{2: "vb kc", 3: "kwkw"}
	for aaak, aaav := range aaa {
		println("aaa", aaak, aaav)
	}
	println("aaa长度和容量", len(aaa), cap(aaa)) // 支持初始化指定索引赋值，且容量就等于大小

	// 注意，数据和切片都支持 len 和 cap 函数，append 是切片专用函数

	// append 只对切片追加元素
	//追加一个元素
	//slice2:=append(slice1,"f")
	//多加多个元素
	//slice2:=append(slice1,"f","g")
	//追加另一个切片
	//slice2:=append(slice1,slice...)
}

// 小结：
/*
1.切片是数组的抽象
2.两者都是索引，且是顺序读写
3.两者都支持角标截取 array[star:end]和slice[star:end]，返回的都是新的切片，数组的切片容量等于原数组大小，切片的新切片容量等于原切片的初始容量或翻倍容量，截取后的新切片索引重置
4.append是切片专用函数，切片和数组都支持 len()、cap() 函数
5.数组的返回类型为 [size]type，切片的返回类型为 []type
6.[]int{}定义的是空切片，因为如果是数组，长度为零的数组没有意义
7.重点：make() 实例化函数，只用来操作 slice, map, or chan (only)
8.两者都支持初始化按照索引指定的值进行赋值，数组大小等于容量，切片容量等于最大角标+1
9.切片和数组，都支持在容易范围内使用角标 array[i] 或 slice[i] 操作，但都不允许超出容量，切片可以自由拓展容量，但需要使用 append() 函数才行！！！

切片的创建方式：
1.直接声明 var s []int
2.初始化 var s = []int{}
3.从数组截取 s := array[star:end]
4.从其它切片截取 s: = slice[star:end]
5.使用 make([]int, l, c)

对比后发现，
数组和切片本就是一家，数组限制较严格，而切片比较灵活，大多部场景切片更合适...
在 Go 语言开发中，切片是使用最多的，尤其是作为函数的参数时，相比数组，通常会优先选择切片，因为它高效，内存占用小
原则：能使用数组的，尽量用切片替代，如果键不是整型才会选择map
*/
