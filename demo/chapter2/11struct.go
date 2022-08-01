package chapter2

import (
	"fmt"
	"os"
)

// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合
// 结构体是开发人员自定义类型集合，是一种新类型的创建方式，跟数组一样，数组是同类型数据的集合，而结构体则完全开放了类型限制（像极了其它语言的类的定义）
// 结构体中的属性和关联方法，叫成员属性和成员方法（即，所属结构体的变量才叫属性，所属结构体的函数才叫方法）

type Circle struct {
	radius float64
}

// 该 method 属于 Circle 类型对象中的方法
// 在结构体中“定义方法”，取决于函数是否接受了指定的结构体：
// 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针，所有给定类型的方法属于该类型的方法集
// 命名类型？结构体类型？
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func Example11() {
	println("Example11:")
	// 结构体的”实例化“直接声明即可
	var c1 Circle // 声明一个 Circle 类型即可！！！【在 Go 语言中，所有的语法组织，理解都趋于扁平化，Circle 类型同 int 类型一样的用法，其它语言可能要 new Circle。。。。】
	c1.radius = 20
	fmt.Println("c1的面积：", c1.getArea())

	c2 := Circle{25}
	fmt.Println("c2的面积：", c2.getArea())

	c3 := &Circle{radius: 30}
	fmt.Println("c3的面积：", c3.getArea())

	// 结构体简单的使用
	type Address struct {
		Province    string
		City        string
		ZipCode     int
		PhoneNumber string
	}
	addr := Address{ // 跟数组的使用完全一样
		"四川",
		"成都",
		610000,
		"0",
	}

	fmt.Printf("%T", addr)
	os.Exit(20000)

	fmt.Println(addr) // {四川 成都 610000 0} // 这种格式，想起了 js 的对象

	// 结构体的引用初始化
	type People struct {
		name  string
		child *People
	}
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	println(relation)

	// 匿名结构体
	// 匿名结构体的类型名是结构体包含字段成员的详细描述，匿名结构体在使用时需要重新定义，造成大量重复的代码，因此开发中较少使用
	msg := &struct {
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)
}

// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("匿名结构体（跟匿名函数很像）：%T\n", msg)
}

/*

关于结构体和类，面向对象的对比理解
Go 没有面向对象，而我们知道常见的 Java、C++ 等语言中，实现类的方法做法都是编译器隐式的给函数加一个 this 指针
而在 Go 里，这个 this 指针需要明确的申明出来，其实和其它 OO 语言并没有很大的区别
C++ 是这样写的：
class Circle {
  public:
    float getArea() {
       return 3.14 * radius * radius;
    }
  private:
    float radius;
}
// 其中 getArea 经过编译器处理大致变为
float getArea(Circle *const c) { // 可以看到，不管是 C++ 和 Go 本质是一样的，形式可能不太一样，而且更多的在于他们的编译器到底为他们做了多少（甚至可以说，一个语言的使用体验，取决于如何定义编译器）
  // ...
}

Go 代码是这样写的：
func (c Circle) getArea() float64 {
    // c.radius 即为 Circle 类型对象中的属性（在 Go 中，结构体的属性是需要显示指定的）
    return 3.14 * c.radius * c.radius // 这里相当于 this.radius，但又由于 Go 的方法是定义在结构体之外的，所以使用 this 显得不易理解。。。。。
}
*/

// 结构体是类型组合体 + 功能体？
// 如此看的话，它跟类没有什么两样了....而结构体的定义显得更加抽象
