package chapter3

// module github.com/jorry2008/goproject-demo
import (
	chapter3 "github.com/jorry2008/goproject-demo/chapter3/demo1" // 别名
	"github.com/jorry2008/goproject-demo/chapter3/demo2"
	"github.com/jorry2008/goproject-demo/chapter5"
	"github.com/jorry2008/goproject-demo/pk1"
)

// 四种引用方案：
import "fmt" // 标准引用

// import F "fmt" // 别名引用，有效解决同名问题（同名公有 结构体、方法名、变量、常量 等）
// import . "fmt" // 省略引用，相当于把 fmt 包直接合并到当前程序中，当前程序可直接使用 Println() 方法
// import _ "fmt" // 匿名引用，只引入并执行包的 init() 函数（一般构建设计模式会用到）

// 1.包：包（package）是多个 Go 源码的集合，包是组织源代码的最小管理单元，是一种高级的代码复用方案。
// - 每个 Go 文件必须所属一个包；
// - 包名是任意的、Go 源文件名也是任意的，两个命名不需要与任何已知的元素有关联；
// - 同一个文件夹下（同级），只能有一种包名，即：一个文件夹一定是一个包，而多个文件夹可以使用一个包名；
// - 包可以定义在任意深度的目录中，包名的定义是不包括目录路径的，但是包在引用时一般使用全路径引用；

// 2.建议：
// - 推荐包名和所在目录同名，结构清晰可读性好
// - 推荐一个文件夹或其子文件夹都应该只包含一个包的源文件
// - 不推荐将一个包的源文件分布在多个不同文件夹下

// 3.小结：包在代码设计时是一种逻辑划分，而在使用包时确是以目录的形态引入！这是一种非常牛逼的理念。
// 在 Go 包中 常量、变量、函数 都是一等公民（其它oop语言唯一的一等公民只有类，其它的所有内容都必须所属在类中）

func Example23() {
	fmt.Println("Example23:")

	// 1.包内调用公共结构体
	resut1 := S1{Name: "aaa"}
	fmt.Printf("%#v \n", resut1)

	// 2.包内调用私有结构体
	resut2 := s1{name: "a"}
	fmt.Printf("%#v \n", resut2)

	// 3.基于包引入，调用其它包的公共内容
	dog1 := pk1.Dog{Name: "jorry", Age: 9, Other: map[string]string{
		"a": "a",
		"b": "b",
	}}
	fmt.Println(dog1)

	// 4.无法调用其它包的私有内容
	//dog2 := pk1.dog{name: "jorry", age: 9, other: map[string]string{
	//	"a": "a",
	//	"b": "b",
	//}}
	//fmt.Printf("%#v", dog2)

	fish1 := demo2.Fish{Name: "fish_name"}
	fmt.Printf("%#v \n", fish1)

	pig1 := chapter3.Pig{Name: "pig_name"}
	fmt.Printf("%#v \n", pig1)

	ch5p := chapter5.Pig{}
	fmt.Printf("%#v \n", ch5p)

}
