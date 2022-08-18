package chapter5

//很多熟悉Go的程序员们都会说到Go是一门很简单的语言，话虽如此，但实际上Go的简单是基于复杂底层的极简包装。
//Go在很多地方均做了“隐式”的转换，这也就导致了很多迷惑点，本文总结了Go开发中几个令人迷惑的地方，如有不当之处请指正。

//nil究竟是什么
//首先明确一点：nil是值而非类型。nil值只能赋值给slice、map、chan、interface和指针。
//
//在Go中，任何类型都会有一个初始值。数值类型的初始值为0，slice、map、chan、interface和指针类型的初始值为nil，对于nil值的变量，我们可以简化理解为初始状态变量。
//
//但nil在实际使用过程中，仍有不少令人迷惑的地方。
//
//var err error
//e := &err
//if e != nil {
//fmt.Printf("&err is not nil:%p\n", e)
//}
//// 输出：&err is not nil:0xc0000301f0
//err是一个接口类型的变量，其初始值为nil，然后对err进行取址操作会发现能成功取到地址，这就是Go和C++最大的不同之一。有C++基础的人在刚接触Go的时候，自然而然的会认为nil是个空指针类型值，上面的代码力证在Go中，nil只是一个表示初始状态的值。
//
//对于slice、map、chan、interface，当值为nil时，不具备可写性。
//
//// 1
//var s []int
//fmt.Printf("%v\n", s[0])
//// 输出panic
//
//// 2
//var c chan int
//val := <-c
//fmt.Printf("%v\n", val)
//// 输出panic
//
//// 3
//var m map[int]int
//m[1] = 123
//// 输出panic
//上面3段代码均会出现panic，对于slice、map、chan类型的nil值变量，可以理解为可读不可写，只有通过make(new)创建的对象实例满足可写性。
//
//参考：https://zhuanlan.zhihu.com/p/105554073
