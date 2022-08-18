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

func Example5_10() {
	fmt.Println("Example5_10:")

	// 断言
	// type assertion 和 type switch 是接口类型断言，是把一个接口类型，转换为具体类型的方法。
	// 很多人会误以为具体类型也可以用这两种方式转换类型，这是错误的。
	// 可以断言的情况：
	// 1.由接口断言到结构体
	// 2.由父接口断言到子接口
	// 指针和非指针实现的方法，断言时的写法不同，语法为： type, ok := xxxx.(T)  // 返回参数 type 为断言之后的类型 T 值，失败则是 nil；参数2 ok 为是否断言成功返回布尔值；
	// 重点：如果类型本身就是断言的类型，则断言成功，会”转换“成这个类型并返回（这就是断言），会将预期的类型返回，从而实现转换

	var p1 A
	p1 = B{}
	v1, ok1 := p1.(B)
	fmt.Println(v1, ok1)

	var p2 A
	p2 = &B{}
	v2, ok2 := p2.(*B)
	fmt.Println(v2, ok2)

	//var a interface{}
	//value, ok := a.(string)
	//if !ok {
	//	fmt.Println("It's not ok for type string")
	//	return
	//}
	//fmt.Println("The value is ", value)

	var t interface{}
	t = 5
	switch v := t.(type) {
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
	//os.Exit(0)

	//var p1 IMale = Son2{} // 匹配 func (s Son2) haha() 方法
	////p1.(Son2).haha()
	//s1, ok := p1.(Son2)
	//if ok {
	//	s1.haha()
	//}

	//var p2 IMale = &Son2{} // ???
	//p1.(Son2).haha()
	//s2, ok := p2.(Son2)
	//if ok {
	//	s2.haha()
	//}

	//var p3 IMale = Son1{} // 匹配 func (s Son1) haha() 方法并实现 IMale 接口
	//p3.(Son1).haha()
	//s3, ok := p3.(Son1)
	//if ok {
	//	s3.say() // 调用了非 IMale 接口方法
	//}
	//
	//var p4 IMale = &Son1{} // 匹配 func (s *Son1) say() 方法并实现 IPeople 接口
	//p4.(*Son1).haha()
	//s4, ok := p4.(*Son1)
	//if ok {
	//	s4.say()
	//}

	// 由父接口断言到子接口【这里要重点理解？？？？？？？？？？？】
	//var iBase IPeople = &Son1{} // son1实现了两个接口 iBase 和 IParent，因为接口本身就是指针，所以它可以直接接受指针
	//iBase.(IMale).haha()
	//
	//var iParent IPeople = &Son1{}
	//iParent.(IPeople).say()

}

//type IPeople interface { // 人
//	say()
//}
//type IMale interface { // 男人
//	haha()
//}
//
//type Son1 struct{}
//
//func (s *Son1) say() {
//	fmt.Println("son1: say")
//}
//func (s Son1) haha() {
//	fmt.Println("son1: haha")
//}
//
//type Son2 struct{}
//
//func (s Son2) haha() {
//	fmt.Println("son2: hi")
//}

// 参考：https://learnku.com/articles/42797
