package chapter2

import "fmt"

// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合
// 结构体跟类一样，也是一种开发人员自定义类型集合，是一种新类型的创建方式

// 结构体中的属性和方法，叫成员属性和成员方法（所属结构体的函数才叫方法，所属结构体的变量才叫属性）

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

func Example0069() {
	// 结构体的”实例化“直接声明即可
	var c1 Circle // 声明一个 Circle 类型即可！！！【在 Go 语言中，所有的语法组织，理解都趋于扁平化，Circle 类型同 int 类型一样的用法，其它语言可能要 new Circle。。。。】
	c1.radius = 20
	fmt.Println(c1.getArea())
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
