package chapter1

import (
	"fmt"
	"unsafe"
)

// VAR_GLASS
// 常量是一个超级简单值的标识符，唯一的特性就是：在程序运行时，不会被修改的量（只读变量）
// 常量中的数据类型只可以是 布尔型、数字型（整数型、浮点型和复数）这种非引用的值类型（所以，常量的类型，仍然属于值类型，可以参与计算）
// 字符串类型除外，它是引用类型，但可以用于常量中
// 与变量相比，差异：常量声明了必须要初始化赋值！表达式赋值必须使用内置函数！（常量后接的是表达式）
const VAR_GLASS = 52  // 首字母有大写，是公共常量
const Var_glass = 25  // 大小写不同，常量就不同，严格区分大小写
const var_glass = 250 // 首字母小写，是私有常量，外部不允许访问

// 枚举定义
const (
	Unknown = 0 // 未知
	Female  = 1 // 女性
	Male    = 2 // 男性
)

const cc string = "abc" // 显式类型定义
const cc1 = "abc"       // 隐式类型定义

// 表达式声明，这个特性限制很大
// 常量可以用 len(), cap(), unsafe.Sizeof() 函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
const (
	a1 = "abc"
	b1 = len(a1)
	c1 = unsafe.Sizeof(a1)
)

// 在枚举常量中，iota 好比一个隐形的计数器
// iota，它本身就是一个常量，只是比较特殊，可以认为是一个可以被编译器修改的常量
// iota 在 const 关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
const (
	i1 = iota // 0
	i2        // 1
	i3        // 2
	i4 = "ha" // 独立值，iota += 1
	i5        // "ha"   iota += 1
	i6 = 100  // iota +=1
	i7        // 100  iota +=1
	i8 = iota // 7,恢复计数【重点在这，必须手动恢复，否则后面的所有值将与最后一个赋值保持一样，但也影响不了 iota 的计数】
	i9        // 8
)

func Example3() {
	fmt.Println("Example3:")

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
	fmt.Println(i1, i2, i3, i4, i5, i6, i7, i8, i9) // 0 1 2 ha ha 100 100 7 8
}
