package chapter5

import "fmt"

// 这里会涉及到 断言、接收器值或指针、隐式转换 等，都跟这个有关

// 为什么要寻址（Addressable），什么叫寻址？为可寻址，就表示无法继续进行链式操作？

// 参考：https://juejin.cn/post/7077800061417029640

//不可寻址
//常量的值。
//基本类型值的字面量。
//算术操作的结果值。
//对各种字面量的索引表达式和切片表达式的结果值。不过有一个例外，对切片字面量的索引结果值却是可寻址的。
//对字符串变量的索引表达式和切片表达式的结果值。
//对字典变量的索引表达式的结果值。
//函数字面量和方法字面量，以及对它们的调用表达式的结果值。
//结构体字面量的字段值，也就是对结构体字面量的选择表达式的结果值。
//类型转换表达式的结果值。
//类型断言表达式的结果值。
//接收表达式的结果值。

// 理解可寻址：https://www.51cto.com/article/684755.html
// 理解二：http://www.wu.run/2021/11/12/not-addressable-in-golang/

type Employee struct {
	Name string
}

func (e Employee) Hi() {
	fmt.Printf("Hi! I am %s.\n", e.Name)
}

func (e *Employee) Hello() {
	fmt.Printf("Hello! I am %s.\n", e.Name)
}

func Example5_12() {
	fmt.Println("Example5_12:")
	var a Employee = Employee{"Alice"}
	a.Hi()
	a.Hello()

	var b interface{} = Employee{"Bob"}
	b.(Employee).Hi()
	// 以下编译报错
	//b.(Employee).Hello() // cannot call pointer method Hello on Employee
	// 可以顺利通过
	dd := b.(Employee)
	dd.Hello() // 可寻址，自动类型转换

	// 过于经典，不便展示^_^
	var c interface{} = &Employee{"Chris"} // new(Employee)
	c.(*Employee).Hi()
	c.(*Employee).Hello()
}

// 编译错误分析：
/*
原因：要调用指针接收器方法，必须要一个指向该值的指针。
这个指针可以是显式的也可以是隐式的，但在形式上必须具有一个可寻址变量的指针。因此，b.(Employee)本身是不可寻址的，需要一个新的指针指向它，直到新指针可寻址为止。
而并非Go中的所有内容都是可寻址的：比较常见的就是接口中的值(如上述实例)和映射条目。
这些情况下，您必须"将无法寻址的值从 接口/映射 中复制出去"以获取可寻址的值，从中可以获取指针并调用指针接收器方法。
这也说明，隐式类型转换对于不可寻址的内容，同样无法实现转换，比如 b.(Employee).Hello() 调用了一个指针接收器方法 .Hello()，此时 b.(Employee) 应该要隐式转换成指针类型，但 b.(Employee) 不可寻址，这个过程发生不了...

// 以上问题，产生于断言的返回值：
类型断言 b.(Employee) 的值是类型 Employee。
方法调用 b.(Employee).Hello() 是 (&b.(Employee)).Hello() 的简写，是因为 func (e *Employee) Hello() 具有指针接收器。
但是，b.(Employee)(一个表达式)不可寻址。 因此，报错
*/

// 对于 T 类型的操作数 x，地址操作 &x 生成 *T 类型的指向 x 的指针。
// 操作数必须是可寻址的，即变量、指针间接或切片索引操作；
// 或可寻址结构操作数的字段选择器；
// 或可寻址数组的数组索引操作。作为可寻址性要求的一个例外，x 也可以是（可能带括号的）复合文字。

// 参考：https://www.codenong.com/43883502/
