package chapter5

import (
	"fmt"
)

// go 存在 4 种类型转换分别为：断言、强制、显式、隐式
// 断言 常见类型转换、强制 在日常不会使用到、显示 是基本的类型转换、隐式 不会注意到
// 断言、强制、显式 三类在 go 语法描述中均有说明，隐式是在日常使用过程中总结出来

type A interface {
	aa()
	bb()
}

type B struct{}

func (b B) aa() {
	fmt.Println("打印...aa")
}

func (b B) bb() {
	fmt.Println("打印...bb")
}

type Animal interface {
	Barking() // 接口方法，只能对指针解引用！
}

type Dog struct{}

func (d Dog) Barking() {
	fmt.Printf("W~W~W~\n")
}

type Pig struct{}

func (c *Pig) Barking() {
	fmt.Printf("En~En~~\n")
}

func AnimalBarking(a Animal) {
	// 这个 a 经过接口处理过了（原来：&Dog{}、Dog{}、Pig{}、&Pig{}都可以调用，现在有些差异，原因是什么？）
	a.Barking() // 参数这里没有类型转化，但是在调用方法时，仍然遵守类型转化
}

func Example5_10() {
	fmt.Println("Example5_10:")

	// 一、以上四个操作说明：调用与方法实现无关，调用者有调用者的自由，方法实现者也有自己的自由，两者之间得益于隐式类型自动转换（本质上调用者类型与接收器类型是一致的，只是发生了自动转换）
	// 即：我们定义普通方法（非接口实现方法），指针接收器和值接收器之间就没有区别，直接忽略，按照实际需求写就可以了。
	d1 := &Dog{}
	d1.Barking() // 隐式自动转换为对象实例

	d2 := Dog{}
	d2.Barking() // 正常调用

	p1 := Pig{}
	p1.Barking() // 隐式自动转换为指针对象实例

	p2 := &Pig{}
	p2.Barking()

	// 二、接口方法的接收器类型问题
	// 为了实现接口而定义方法时，指针接收器和值接收器之间就有区别了：
	println("-------")
	// 2.1.接口方法实现中，包含的全部是值类型接收器，就可以任意转换为接口类型
	d11 := &Dog{}
	AnimalBarking(d11)

	d22 := Dog{}
	AnimalBarking(d22)

	// 2.2.接口方法实现中，至少包含一项指针类型接口器，只能将指针类型传递给接口类型
	d33 := &Pig{}
	AnimalBarking(d33)

	//d44 := Pig{}
	//AnimalBarking(d44) // Pig{}值类型未实现 Animal 接口，&Pig{}指针类型实现了 Animal 接口，其方法是 func (c *Pig) Barking()，指针类型接收器

	// 三、断言
	// type assertion 和 type switch 是接口类型断言，是把一个接口类型，转换为具体类型的方法。
	// 很多人会误以为具体类型也可以用这两种方式转换类型，这是错误的。
	// 可以断言的情况：
	// 1.由接口断言到结构体
	// 2.由父接口断言到子接口
	// 指针和非指针实现的方法，断言时的写法不同，语法为： type, ok := xxxx.(T)  // 返回参数 type 为断言之后的类型 T 值，失败则是 nil；参数2 ok 为是否断言成功返回布尔值；
	// 重点：如果类型本身就是断言的类型，则断言成功，会”转换“成这个类型并返回（这就是断言），会将预期的类型返回，从而实现转换
	println("-------")
	var pz1 A
	pz1 = B{}
	v1, ok1 := pz1.(B)
	fmt.Printf("断言：%T %v \n", v1, ok1)

	var pz2 A
	pz2 = &B{}
	v2, ok2 := pz2.(*B)
	fmt.Printf("断言：%T %v \n", v2, ok2)

	var a interface{}
	a = "str test"
	value, ok := a.(string)
	if !ok {
		fmt.Println("It's not ok for type string")
		return
	}
	fmt.Println("The value is ", value)

	var t interface{}
	t = "abc"
	switch v := t.(type) { // 这是一种固定的语法结构，不可拆解
	default:
		fmt.Printf("unexpected type %T", v) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", v) // t has type bool
	case int:
		fmt.Printf("integer %d\n", v) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *v) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *v) // t has type *int
	}

	// 四、从这个问题，引导出可寻址、不可寻址
	var pp1 IMale = Son2{} // 匹配 func (s Son2) haha() 方法
	//p1.(Son2).haha()
	s1, ok := pp1.(Son2)
	if ok {
		s1.haha()
	}

	var pp2 IMale = &Son2{} // ???
	pp1.(Son2).haha()
	s2, ok := pp2.(Son2)
	if ok {
		s2.haha()
	}

	var pp3 IMale = Son1{} // 匹配 func (s Son1) haha() 方法并实现 IMale 接口
	pp3.(Son1).haha()
	s3, ok := pp3.(Son1)
	if ok {
		s3.say() // 调用了非 IMale 接口方法
	}

	var pp4 IMale = &Son1{} // 匹配 func (s *Son1) say() 方法并实现 IPeople 接口
	pp4.(*Son1).haha()
	s4, ok := pp4.(*Son1)
	if ok {
		s4.say()
	}

	// 由父接口断言到子接口【这里要重点理解？？？？？？？？？？？】
	//var iBase IPeople = &Son1{} // son1实现了两个接口 iBase 和 IParent，因为接口本身就是指针，所以它可以直接接受指针
	//iBase.(IMale).haha()
	//
	//var iParent IPeople = &Son1{}
	//iParent.(IPeople).say()

}

type IPeople interface { // 人
	say()
}
type IMale interface { // 男人
	haha()
}

type Son1 struct{}

func (s *Son1) say() {
	fmt.Println("son1: say")
}
func (s Son1) haha() {
	fmt.Println("son1: haha")
}

type Son2 struct{}

func (s Son2) haha() {
	fmt.Println("son2: hi")
}

// 参考：https://learnku.com/articles/42797
