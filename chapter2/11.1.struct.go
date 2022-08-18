package chapter2

import (
	"fmt"
	"reflect"
)

// 在 go 中，万物皆类型！
// new(Type) 和 &Type{} 等价

type CustomName struct{}
type TypeAlias = map[int]map[string]string // 类型别名：TypeAlias 只是右边类型的别名（本质上两者是对同一个类型的多种称呼）(赋值符号两边完全一样，表示同一个东西)
type CustomType map[int]map[string]string
type CustomType1 []int
type cb func(s string)                  // 看吧，简单的回调函数类型
type cb1 func(s1 CustomType) (s string) // 在已有类型之上，定义新类型！这个用处太方便了，可以将长长的类型定义包裹起来使用（定义一个函数类型）

// 以下是官方定义，可以看到，byte 和 rune 并不是一个新类型，而是方便识别的别名
//type byte = uint8
//type rune = int32

// 结构体的深入理解
type Cat struct {
	Name    string
	Age     int
	Color   string
	Teacher *string // 接受字符串指针，就是说明必须是引用传递，方法内的修改反映到方法外
}

type Cat2 struct {
	// 在定义结构体字段时，除字段名称和数据类型外，还可以使用反引号为结构体字段声明元信息，这种元信息称为Tag，用于编译阶段关联到字段当中
	// `json:"name"` 指定序列化后的 Name =》name
	Name  string `json:"name"` // 序列化转化
	Age   int
	Color string
}

// 一个独立的方法
func setCat(cat *Cat) {
	cat.Name = "小小"
	fmt.Println(cat.Name)
}

func Example11_2() {

	// 定义一个类型 VS 创建一个类型别名
	type mapSS map[int]map[string]string    // 直接定义一个新的类型叫 mapSS（mapSS 和 map[int]map[string]string 如同克隆）
	type mapSS1 = map[int]map[string]string // 给 map[int]map[string]string 类型取了一个新的名称，本质上是一个指针（且这里定义的是一个内置类型别名，不能对其进行拓展）
	// type的用法，参考：https://juejin.cn/post/6844903926450372616

	// 方式1
	students := make(map[int]map[string]string) // map的嵌套结构
	students[1] = map[string]string{
		"name": "张三",
		"sex":  "男",
	}
	students[2] = map[string]string{
		"name": "张四",
		"sex":  "男",
	}

	fmt.Println(students) // 输出 =》map[1:map[name:张三 sex:男] 2:map[name:张四 sex:男]]

	var varmapSS mapSS
	varmapSS = students
	fmt.Println(varmapSS) // 输出 =》map[1:map[name:张三 sex:男] 2:map[name:张四 sex:男]]

	var varmapSS1 mapSS1
	varmapSS1 = students
	fmt.Println(varmapSS1) // 输出 =》map[1:map[name:张三 sex:男] 2:map[name:张四 sex:男]]

	// 结论：原生定义的类型变量、自定义type类型变量、类型别名，三者创建的类型，在使用上是一致的

	// 方式2
	stus := mapSS{
		1: {
			"name": "王武",
			"age":  "20",
		},
		2: {
			"name": "王六",
			"age":  "24",
		},
		3: {
			"name": "王七",
			"age":  "30",
		},
	}
	fmt.Println(stus)
	a, b := stus[1]["name"]
	fmt.Println(a, reflect.TypeOf(b))
	// 输出 =》 map[1:map[age:20 name:王武] 2:map[age:24 name:王六]]

	// 方式1
	var cat Cat
	fmt.Println(cat) // { 0  <nil>}

	// 方式2
	cat2 := Cat{}     // 与 &Cat{} 写法有什么差别？
	fmt.Println(cat2) // { 0  <nil>}

	// 方式3
	// new一个指针，引用类型
	cat3 := new(Cat) //等价 （var cat3 *Cat = new(Cat)）
	(*cat3).Name = "aaa"
	cat3.Name = "bbb"  // 自动解引用，编译时自动转化为 (*cat3).Name = "bbb"
	fmt.Println(*cat3) // {bbb 0  <nil>}

	// 方式4
	cat4 := &Cat{} // new(Type) 和 &Type{} 等价
	(*cat4).Name = "111"
	fmt.Println(*cat4) // {111 0  <nil>}

	// 结构体切片
	str := ""
	cat15 := Cat{"小花", 10, "红色", &str}
	cat25 := Cat{"小白", 1, "白色", &str}
	sliceCat5 := []Cat{cat15, cat25} // 习惯这些结构
	println(sliceCat5)

	// 结构体传递
	cat6 := new(Cat) // 返回值就是 *Type
	(*cat6).Name = "大大"
	// 此处参数是指针，所以是引用传递，会修改当前对象
	setCat(cat6) // 这里的结构体，本身就是指针，直接传过去就可以了，为什么？
	fmt.Println(cat6)
	// 输出 =》
	// { "小小" }
	// { "小小" }

	// map判断是否存在key
	// _, ok := map1[1]

}

// 3种类型来源对应不同的使用行为

// 1.自定义结构体类型
func (s CustomName) funcName1() { // 这个 CustomName 类型的变量 s 称为方法接收器，专门用来传递结构体的，相当于其它语言中的 this

}

// 2.自定义类型，与普通定义的结构体类型使用上完全一致
func (s CustomType) funcName2() {

}

// 3.基于内建类型的别名，不允许对其进行修改！！！
//func (s TypeAlias) funcName3() {
//
//}
