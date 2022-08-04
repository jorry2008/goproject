package chapter2

import "fmt"

// 类型转换：自动转换，强制转换，精度问题
// 整型、布尔、浮点、字符串、虚数 基本类型的转换

func Example16() {
	println("Example16:")

	// 1.自动转换
	fmt.Printf("%T \n", 25.2+10) // float64，自动转化，保住了精度

	// 2.强制转换
	// 类型转换用于将一种数据类型的变量转换为另外一种类型的变量，
	// type_name(expression)
	// 布尔类型，与整型之间不能转化

	var a int64 = 3
	var b int32
	b = int32(a) // 强制转换为 int32，才能进行同类型赋值操作
	fmt.Printf("b 为 : %d \n", b)

	varint := 25
	varfloat := 25.2
	varstring := "测试"

	println("整型25转为浮点：", float32(varint))
	println("浮点转化为整型 ：", int(varfloat)) // 会有精度损失

	for sk, sv := range []byte(varstring) {
		println(sk, sv)
	}
	bytess := []byte(varstring)
	fmt.Printf("%q", bytess[:3]) // 返回：测

	// 3.类型断言
	// 指针和非指针实现的方法，断言时的写法不同，语法为： type, ok := xxxx.(T)  // 返回参数 1 type 为断言之后的类型值失败则是 nil，参数2 ok 为是否断言成功返回布尔值
	// 重点：如果类型本身就是断言的类型，则断言成功，会”转换“成这个类型并返回（这就是断言），会将真实的类型返回，从而实现转换
	// 可以断言的情况：
	// 1.由接口断言到结构体
	// 2.由父接口断言到子接口
	var p1 IParent = &Son1{} // 指针实现的方法hi
	// p1.(*Son1).hi() // 直接使用，返回的就是转化后的类型
	son, ok := p1.(*Son1)
	if ok {
		son.hi()
	}

	var p2 IParent = Son2{} // 非指针实现的方法hi【注意：这里可以直接赋值给接口类型，原因在于？？？？？？？？？？？？？？？？？？】
	//p2.(Son2).hi()
	son_, ok := p2.(Son2)
	if ok {
		son_.hi()
	}

	// 由父接口断言到子接口【这里要重点理解？？？？？？？？？？？】
	var iBase IBase = &Son1{} // son1实现了两个接口 iBase 和 IParent，因为接口本身就是指针，所以它可以直接接受指针
	iBase.(IParent).hi()

	var iParent IBase = &Son1{}
	iParent.(IBase).hello()

	// 4.向上造型
	// 在 go 语言中，将父结构嵌入到子结构，就是结构体组合操作，这个叫继承
	// 想从子结构转型为父级，直接： son.parent 即可，这个过程叫 向上造型
}

// 以下定义的接口的实现类型，他们之间的类型上下转化

type IBase interface {
	hello()
}
type IParent interface {
	hi()
}
type Son1 struct{}

func (son1 *Son1) hi() {
	fmt.Println("son1: hi")
}
func (son1 *Son1) hello() {
	fmt.Println("son1: hello")
}

type Son2 struct{}

func (son2 Son2) hi() {
	fmt.Println("son2: hi")
}

// 参考：https://blog.csdn.net/qq_43413788/article/details/113651012

// 以上遇到的几个问题，本质还是引用和值的问题，受体方法传入的是值还是引用，对整个理解至关重要？？？？？？？？
