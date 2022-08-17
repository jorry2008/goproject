package chapter5

import "fmt"

// 所有类型都是可以绑定方法的！即所有类型都有隐式解引用、隐式取引用...
// go语言的隐式解引用，三种场景下的演示。

type node struct {
	Name string
	Next *node
}

func (n node) Say() {
	n.Name = "值类型"
}

func (n *node) SayPointer() { // n 表示绑定类型的接受器
	(*n).Name = "指针类型" // 这里应该也转化了 (*n).Name ??????
}

type newSlice []int

func (s newSlice) Hello() {
	s[0] = 100
}

func (s *newSlice) HelloPointer() {
	(*s)[0] = 100
}

func Example5_9() {
	fmt.Println("Example5_9:")

	// 一、先来解决第一个问题：结构体为什么要使用指针传递？
	// 结构体使用 var 还是 new() 创建，哪个好？

	//var m1 *Member
	//m1.name = "小明" //错误用法，未初始化,m1为nil
	//
	//m1 = &Member{}
	//m1.name = "小明" //初始化后，结构体指针指向某个结构体地址，才能访问字段，为字段赋值。
	//
	////复制代码另外，使用Go内置new()函数，可以分配内存来初始化结构休，并返回分配的内存指针，因为已经初始化了，所以可以直接访问字段。
	//var m2 = new(Member)
	//m2.name = "小红"

	// 二、结构体类型的自动解引用，取决于访问了复杂类型的属性，指针会解引用为原始值（值类型）
	n5 := new(node) // 等价 var n5 *node = new(node)
	(*n5).Name = "Jerry"
	n5.Name = "Tom" // 两种形式都支持，正常逻辑不应该这么操作，但Go语言自带隐式解引用，这里底层转化为 (*n5).Name = "Tom"
	fmt.Println(*n5)

	// 三、值类型的的 隐式解引用/隐式取引用，取决于接收器的类型（是值还是指针），与类型本身无关（这样做的好处就是，是取值还是取引用直接由接收器类型决定，这样条件单一不容易出错，可高效开发）

	n1 := &node{Name: "hui"} // 结构体是值类型，这里返回指针
	n1.SayPointer()          // 当前是值类型指针，形参和实参相同，不需要自动转化
	fmt.Println(*n1)

	n2 := &node{Name: "hui"}
	n2.Say() // Say() 是值接收器，go编译器为我们加上*自动解引用 func (*n).Say()，这样类型才能保持一致，可以从这里传递到函数接收器（这里是形参，接收器叫实参）
	fmt.Println(*n2)

	n3 := node{Name: "hui"}
	n3.SayPointer() // SayPointer() 是指针接收器，go编译器为我们加上&自动取引用 func (&n).Say()，这样类型才能保持一致，可以从这里传递到函数接收器（这里是形参，接收器叫实参）
	fmt.Println(n3)

	n4 := node{Name: "hui"}
	n4.Say()
	fmt.Println(n4)

	// 四、引用类型的 解引用/取引用，取决于接收器的类型（是引用还是指针），与类型本身无关，引用类型有个特点：不论是自动解还是自动取，都将受到影响
	s1 := make(newSlice, 5)
	s1.HelloPointer() // 引用类型 & 转化 为指针类型
	fmt.Println(s1)

	s4 := make(newSlice, 5)
	s4.Hello() // 本身就为引用类型，受影响，不需转化
	fmt.Println(s4)

	var s2 *newSlice
	s3 := make(newSlice, 5)
	s2 = &s3
	s2.Hello() // 指针类型 * 转换 为引用类型
	fmt.Println(*s2)

	var s5 *newSlice
	s6 := make(newSlice, 5)
	s5 = &s6
	s5.HelloPointer() // 本身就是指针类型，受影响，不需要转化
	fmt.Println(*s5)
}
