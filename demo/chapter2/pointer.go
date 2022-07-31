package chapter2

// 在Go语言中的值类型有：int、float、bool、string、array、struct 等
// 它们都有对应的指针类型，比如：*int、*float 等

//对变量进行取地址（&）操作，可以获得这个变量的指针变量
//对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值
//指针变量的值是变量的内存地址

func Example009() {

	println("Pointer:")

	x := 9
	y := 10
	//fmt.Printf("%T", &x) // 指针类型 *int

	*&x = 25 // 引用和解引用，x还是本身，所以这里重新对 x 赋值了

	//z := &y // 简写
	var z *int = &y // z为引用类型，才能接受 &y 引用
	*z = 100        // 函数传引用，然后函数体内部 *z 操作，对原实参产生影响，*z 表示的就是原来最初的那个 y

	println(x) // 25
	println(z) // 0xc000109f40
	println(y) // 100

}
