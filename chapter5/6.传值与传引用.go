package chapter5

import (
	"fmt"
)

// 很多人无法清晰的知道 传值、传引用？
// 一方面，自己有太多的执念，将简单的内容复杂化了。
// 另一方面，基础不扎实，特别是数据类型的相关概念没有理解透。

// 值类型：int、float、bool、array、sturct
// 引用类型：slice、map、channel、interface、func、（特殊）string
// 指针类型：独立体系（根据它的定义来就可以了）

// go只有传值与php一样？？？？

type Cat struct {
	Name    string                     // 默认值 ""
	Age     int                        // 默认值 0
	Color   string                     // 默认值 ""
	Teacher *string                    // 默认值 nil
	cb      func(map2 *map[string]int) // 默认值 nil // 回调函数
}

func Reference() {
	fmt.Println("Reference:")

	// 自定义3个数据类型（只声明未初始化）
	var varInt int
	var varSlice []int
	var varCat Cat // 自定义类型同样有零值（即每个属性自身的默认值），这里的结构体用法就是一个普通值类型的用法，完全一样

	fmt.Println(varInt)
	fmt.Println(varSlice)
	fmt.Println(varCat)

	SetInt(varInt)
	fmt.Println(varInt)

	varSlice = append(varSlice, 8)
	fmt.Printf("%T \n", varSlice) // 输出 []int，slice 本身就是引用类型，不需要转化为指针（也可以转化）
	SetSlice(varSlice)            // varSlice 本身是引用类型
	fmt.Println(varSlice)

	fmt.Printf("%T \n", varCat)
	SetCat(varCat)
	fmt.Println(varCat) // 值仍然未变，它表现出来的行为跟普通类型一样，即不是引用，也不是指针

	SetIntPoint(&varInt) // 参数为指针时，值被修改了
	fmt.Println(varInt)

	varSlice = append(varSlice, 88)
	SetSlicePoint(&varSlice) // 将引用类型，以指针的方式传递！！
	fmt.Println(varSlice)

	SetCatPoint(&varCat)
	fmt.Println(varCat) // 值仍然未变，它表现出来的行为跟普通类型一样

	// 小节：自定义结构体类型的用法与普通类型，完全一样！

	//// 六、结构体值类型
	//// new一个指针，引用类型
	//cat3 := new(Cat) //等价 （var cat3 *Cat = new(Cat)）
	//(*cat3).Name = "aaa"
	//cat3.Name = "bbb"  // 两种写法都可以，底层转化为=》(*cat3).Name = "bbb"
	//fmt.Println(*cat3) // 输出 =》{bbb 0  <nil>}
	//fmt.Println(cat3)  // 它本身是一个地址值
	//
	//// 方式4
	//cat4 := &Cat{}
	//(*cat4).Name = "111"
	//cat4.Name = "222"  // 底层转化
	//fmt.Println(*cat4) // =》{222 0  <nil>}
	//fmt.Println(cat4)  // 它本身是一个地址值
	//
	//cat5 := Cat{}
	//cat5.Name = "333"
	//SetCatPoint(&cat5) // 传入指针
	//SetCat(cat5)
	////fmt.Println(*cat5) // 没有???
	//fmt.Println(cat5)
}

func SetInt(i int) {
	i = 88
}

func SetIntPoint(i *int) {
	*i = 89
}

func SetSlice(s []int) {
	s[0] = 9
}

func SetSlicePoint(s *[]int) {
	(*s)[0] = 99
}

func SetCat(c Cat) { // 这里是否表示指标（是否带*），取决于传入值本身的类型！！！
	c.Name = "Tom"
}

func SetCatPoint(c *Cat) { // 这里是否表示指标（是否带*），取决于传入值本身的类型！！！
	c.Name = "Tom"
	c.Age = 50
	c.Color = "red"
	//(*c).Name = "Tom" // 底层自动转化，就这一点点区别
}
