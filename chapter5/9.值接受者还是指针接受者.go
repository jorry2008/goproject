// Package chapter5 主要解释Go语言中的自动解引用/自动取引用
package chapter5

import "fmt"

// 所有类型都是可以绑定方法的！
// 以下实例，从go语言的隐式转换（自动解引用/自动取引用）理解接收器的类型问题：即值接受还是指针接受

// 用法小结：
// 1.所有类型都有隐式解引用、隐式取引用行为
// 2.所有类型的接收器类型取决于自身的类型提示，跟调用方法的类型本身没有关系（调用它的可能是值类型、引用类型、指针类型）
// 3.指针访问属性时，也会自动解引用（结构体等？）【这里还缺一些资料】

// 方法接收器的 自动解引用/取引用，意义何在？
/*
类型设计好了，相关的方法也绑定了，开发人员在写方法时，不会考虑其它开发人员如何调用（值调用、引用调用、指针调用），
甚至其它公司的其它团队的其它开发人员，他们的调用形式是什么，是 var s struct 还是 new(struct) 或者 &struct{}，
否则这程序也应该写不出来，或者自己写给自己用，而且用法只能是一种。
实际上，绑定在这些类型上的函数的接收器的意义如同 this 一样，那这个 this 到底是啥类型？应该由设计人员自己决定！
开发设计人员，在设计方法时，使用什么类型方便它就使用哪个类型，对于使用者而言，也随意使用，这中间肯定会导致类型不匹配的问题，
很明显，这个问题是 go 语言自动完成。
这就是自动类型转换（自动解引用、自动取引用）
*/

// 值类型演示
type node struct {
	Name string
	Next *node
}

func (n node) Say() {
	n.Name = "值类型"
}

func (n *node) SayPointer() { // n 表示绑定类型的接受器
	(*n).Name = "指针类型" // 同样，这里也有解引用（如果指针类型指向的原始类型为引用类型时）
	n.Name = "这里也是解引用"
}

// 引用类型演示
type newSlice []int

func (s newSlice) Hello() {
	s[0] = 100
}

func (s *newSlice) HelloPointer() {
	(*s)[0] = 100
}

func Example5_9() {
	fmt.Println("Example5_9:")

	// 一、结构体类型 自动解引用
	n1 := new(node)      // 等价 var n5 *node = new(node) 也等价于 var n5 *node = &node{}
	(*n1).Name = "Jerry" // 手动解引用
	n1.Name = "Tom"      // 自动解引用，这里底层转化为 (*n1).Name = "Tom"
	fmt.Println(*n1)

	// 二、值类型的的 隐式解引用/隐式取引用
	n2 := &node{Name: "hui"} // 结构体是值类型，这里返回指针
	n2.SayPointer()          // 当前是值类型指针，形参和实参相同，不需要自动转化
	fmt.Println(*n2)

	n3 := &node{Name: "hui"}
	n3.Say() // Say() 是值接收器，go编译器为我们加上*自动解引用 func (*n).Say()，这样类型才能保持一致，可以从这里传递到函数接收器（这里是形参，接收器叫实参）
	fmt.Println(*n3)

	n4 := node{Name: "hui"}
	n4.SayPointer() // SayPointer() 是指针接收器，go编译器为我们加上&自动取引用 func (&n).Say()，这样类型才能保持一致，可以从这里传递到函数接收器（这里是形参，接收器叫实参）
	fmt.Println(n4)

	n5 := node{Name: "hui"}
	n5.Say()
	fmt.Println(n5)

	// 三、引用类型的 自动解引用/自动取引用
	s1 := make(newSlice, 5)
	s1.HelloPointer() // 引用类型 & 转化 为指针类型
	fmt.Println(s1)

	s2 := make(newSlice, 5)
	s2.Hello() // 本身就为引用类型，受影响，不需转化
	fmt.Println(s2)

	var s3 *newSlice
	s4 := make(newSlice, 5)
	s3 = &s4
	s3.Hello() // 指针类型 * 转换 为引用类型
	fmt.Println(*s3)

	var s5 *newSlice
	s6 := make(newSlice, 5)
	s5 = &s6
	s5.HelloPointer() // 本身就是指针类型，受影响，不需要转化
	fmt.Println(*s5)
}
