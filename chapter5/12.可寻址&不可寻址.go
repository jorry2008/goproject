package chapter5

import (
	"fmt"
	"unsafe"
)

// 了寻址这个特性，可以解释很多关于 断言、接口、接收器值或指针、隐式转换 等疑惑！

// 可寻址列表：
/*
1.变量：&x // &name //0xc00003c1f0
2.指针：&*x // unsafe.Pointer(&person{"yif"}) //0xc00003c1f0
3.数组元素索引: &a[0] //
4.切片 //
5.切片元素索引：&s[0] //
6.组合字面量: &struct{X type}{value}
*/

// 不可寻址列表：
/*
1.常量的值
2.基本类型值的字面量
3.算术操作的结果值
4.对各种字面量的索引表达式和切片表达式的结果值。不过有一个例外，对切片字面量的索引结果值却是可寻址的
5.对字符串变量的索引表达式和切片表达式的结果值
6.对字典变量的索引表达式的结果值
7.函数字面量和方法字面量，以及对它们的调用表达式的结果值
8.结构体字面量的字段值，也就是对结构体字面量的选择表达式的结果值
9.类型转换表达式的结果值
10.类型断言表达式的结果值
11.接收表达式的结果值
*/

// 参考：https://juejin.cn/post/7077800061417029640

// 理解可寻址：https://www.51cto.com/article/684755.html
// 理解二：http://www.wu.run/2021/11/12/not-addressable-in-golang/

type Profile struct {
	Name string
}

func new1() Profile {
	return Profile{Name: "a"}
}

func (p Profile) ProfileFunc() {
	fmt.Println("Profile Func")
}

func (p *Profile) PtrProfileFunc() {
	fmt.Println("Ptr Profile Func")
}

type Employee struct {
	Name string
}

func (e Employee) Hi() {
	fmt.Printf("Hi! I am %s.\n", e.Name)
}

func (e *Employee) PtrHello() {
	fmt.Printf("PtrHello! I am %s.\n", e.Name)
}

func Example5_12() {
	fmt.Println("Example5_12:")

	// new1() 同 Profile{Name: "a"} 一样，仅仅是结构体字面量 {a} 或 {b}，为不可寻址的值本身
	fmt.Println(new1())              // {a}
	fmt.Println(Profile{Name: "b"})  // {b}
	new1().ProfileFunc()             // 非指针接受器，可直接调用
	Profile{Name: "b"}.ProfileFunc() // 同上
	//new1().PtrProfileFunc()             // 指针接收器，不能由结构体本身进行调用，因为它不是指针，无法寻址
	//Profile{Name: "b"}.PtrProfileFunc() // 同上【其实，在 golang 内部是会自动类型转换的，golang 会自动帮忙取地址操作，本来这句是对的，但在这里好像失效了？？？其实不是的，有一个前提：对于不可寻址的对象无法直接调用，即无法激活类型自动转换】
	//fmt.Println(unsafe.Pointer(Profile{"ha"})) // 无法寻址
	// 但可以这样调用
	pp1 := new1()
	pp1.PtrProfileFunc() // pp1 是变量，可寻址，隐式自动类型转化，可以调用

	fmt.Println(&Profile{Name: "c"})
	(&Profile{Name: "c"}).PtrProfileFunc()
	(&(Profile{Name: "e"})).PtrProfileFunc()
	fmt.Println("获取结构体指针的地址值：", unsafe.Pointer(&Profile{"haha"})) // 可寻址

	var a Employee = Employee{"Alice"}
	a.Hi()
	a.PtrHello() // 隐式自动类型转化

	var b interface{} = Employee{"Bob"}
	b.(Employee).Hi()
	// 以下编译报错
	//b.(Employee).PtrHello() // cannot call pointer method PtrHello on Employee，b.(Employee) 相当于 Employee{"Bob"} 无法寻址，所以这里无法直接调用指针接收器方法
	// 可以顺利通过
	dd := b.(Employee)
	dd.PtrHello() // 使用新变量包裹 Employee{"Bob"} 后，可寻址，自动类型转换

	// 过于经典，不便展示^_^
	var c interface{} = &Employee{"Chris"} // new(Employee)，这种方式也可以将指针直接填充到 map[string]*Employee
	c.(*Employee).Hi()
	c.(*Employee).PtrHello()

	var a1 map[string]*Employee
	a1 = make(map[string]*Employee)
	a1["aaa"] = &Employee{"Chris1"}
	fmt.Println(a1)
	a2 := map[string]*Employee{"aaa": {"Chris2"}} // 这里有类型自动推断
	fmt.Println(a2)
}

// 编译错误分析：
/*
原因：要调用指针接收器方法，必须要一个指向该值的指针。
这个指针可以是显式的也可以是隐式的，但在形式上必须具有一个可寻址变量的指针。因此，b.(Employee)本身是不可寻址的，需要一个新的指针指向它，直到新指针可寻址为止。
而并非Go中的所有内容都是可寻址的：比较常见的就是接口中的值(如上述实例)和映射条目。
这些情况下，您必须"将无法寻址的值从 接口/映射 中复制出去"以获取可寻址的值，从中可以获取指针并调用指针接收器方法。
这也说明，隐式类型转换对于不可寻址的内容，同样无法实现转换，比如 b.(Employee).PtrHello() 调用了一个指针接收器方法 .PtrHello()，此时 b.(Employee) 应该要隐式转换成指针类型，但 b.(Employee) 不可寻址，这个过程发生不了...

// 以上问题，产生于断言的返回值：
类型断言 b.(Employee) 的值是类型 Employee。
方法调用 b.(Employee).PtrHello() 是 (&b.(Employee)).PtrHello() 的简写，是因为 func (e *Employee) PtrHello() 具有指针接收器。
但是，b.(Employee)(一个表达式)不可寻址。 因此，报错
*/
