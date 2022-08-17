package chapter5

import "fmt"

// 引用的历史：
// 1.我们应该先了解引用的初衷：通常意义的引用对于指针而言不容易出错，使用方便，是一种不透明引用，而指针本身是一种透明引用，接近底层容易出错，而引用的初衷是想实现对指针的替代。
// 2.相比于其它语言，go 语言的引用较特殊：go 引用是基于指针类型，且只是针对特定类型的优化（go语言中，没有像其它语言一样的引用，它是对指针进一步包装的结果，即强制某些类型为引用类型）。
// 3.为了方便归纳，在 go 教学中，整个类型体系分为 值类型、引用类型，而引用类型包含指针类型（透明引用），引用类型中的 string 行为比较特殊。
// 4.特点：在其它语言中，对同一个目标的引用，在任何情况下无论怎么操作指向的目标永远是同一个目标，而在 go 语言中，引用更像是一种无法解引用的指针（不能使用 *ptr 操作），且它允许覆盖。（看以下的例子）

// 我自己的理解：
// 1.值类型：bool、int、float、complex、byte、rune、uintptr、array、struct
// 2.引用类型：string、slice、map、func、interface、channel，其中 string 较特殊
// 3.指针类型：pointer （一种透明引用）

// 官方明确，在 go 语言中没有引用类型（这个理解一下就行），我们认为是一种特殊的区别于其它语言的 go 专用引用
// 参考：https://github.com/golang/go/commit/b34f0551387fcf043d65cd7d96a0214956578f94

// 关于为什么go没有引用的例子：
// 如果他们是引用变量，那么下面这段程序将打印 false
func fnuc1(map2 map[int]int) {
	// 使用 make 申请了空间
	map2 = make(map[int]int) // 在引用特性中，这里的实参应该是无法覆盖的，但在go中，这里是可以覆盖的，导致map2重新指向了新的地址，对原map2没有任何影响！！！
}

func Example5_5() {
	fmt.Println("Example5_5:")

	var map1 map[int]int     // 只声明，没有分配空间，零值为 nil，map是go的特殊”引用“
	fmt.Println(map1 == nil) // true
	fnuc1(map1)
	fmt.Println(map1 == nil) // true
}

// 关于 map 解答，引用是一种指针的不透明使用：
// 很多教程都说 go 中的引本质就是由指针实现的，并不是真正的引用，为什么呢？以map举例：
// Go 源代码中显示 https://golang.org/src/runtime/hashmap.go
// map 底层是一个指向 hmap 的指针，这就可以解释即使函数传参是按值传递，由于传递的是指针的拷贝，指针指向的底层 hmap 并没有改变，所以可以在函数内部改变 map
// 因此，map 函数内可以被修改，但它却可以被覆盖（这个就不是引用了）

// 最后一个疑问：基于引用，slice、map、channel、string，可以联动修改值，那么 interface、func 引用特性能干嘛？（两者在底层上当然是提高效率和降低存储）？

// ???
