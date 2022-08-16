package chapter5

import "fmt"

// 我们应该先了解引用的历史：引用对于指针而言不容易出错，使用方便，是一种不透明引用，而指针本身是一种透明引用，接近底层容易出错，所以引用本来是一种对指针的替代方案。
// 而 go 语言对于其它语言而言，它的引用较特殊：go引用是基于指针类型，且只是针对特定类型的优化。

// 引用类型包括：slice、map、channel、interface、func、string （go语言中，没有像其它语言一样的引用，它是对指针进一步包装的结果，且go引用只是针对特定类型的优化！）

// 基于引用，slice、map、channel、string，可以联动修改值，那么 interface、func 引用特性能干嘛？（两者在底层上当然是提高效率和降低存储）

// .....

// 首先明确一下，在 go 语言中，是没有引用类型的，参考：https://github.com/golang/go/commit/b34f0551387fcf043d65cd7d96a0214956578f94
// 但就是有那么一种类型，非常像引用，但又不符合引用的定义，也不是指针类型，所以它到底是什么类型呢？

// go 语言中到底有没有引用？没有的
// Maps 与 Channels 依然不是引用变量。如果他们是引用变量，那么下面这段程序将打印 false.

func fn(m map[int]int) {
	m = make(map[int]int) // 在引用特性中，这里的实参应该是无法覆盖的，但在go中，这里是可以覆盖的，导致m重新指向了新的地址，对原m没有任何影响！！！
}

func Example844564() {
	var m map[int]int
	fn(m)
	fmt.Println(m == nil)
}

// 都说，go 中的引本质就是由指针实现的，并不是真正的引用，为什么呢？以map举例：
// Go 源代码中显示 https://golang.org/src/runtime/hashmap.go
// map 底层是一个指向 hmap 的指针，这就可以解释即使函数传参是按值传递，由于传递的是指针的拷贝，指针指向的底层 hmap 并没有改变，所以可以在函数内部改变 map
// 因此，map 函数内可以被修改，但它却可以被覆盖（这个就不是引用了）

// 最后，探讨一下 func、interface、string 三者的引用特性？？？
