package chapter2

import (
	"fmt"
	"reflect"
)

// 结构体的深入理解
type Cat struct {
	Name    string
	Age     int
	Color   string
	Teacher *string // 接受字符串指针，就是说明必须是引用传递，方法内的修改反映到方法外
}

type Cat2 struct {
	// `json:"name"` 指定序列化后的 Name =》name
	Name  string `json:"name"` // 这是个什么鬼？
	Age   int
	Color string
}

// 一个独立的方法
func setCat(cat *Cat) {
	cat.Name = "小小"
	fmt.Println(cat.Name)
}

func Exaple11_2() {

	// 结构体的变化
	type mapSS map[int]map[string]string // 这是什么用法？定义复合类型？

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
	students[3] = map[string]string{
		"name": "张五",
		"sex":  "女",
	}
	fmt.Println(students)
	// 输出 =》map[1:map[name:张三 sex:男] 2:map[name:张四 sex:男] 3:map[name:张五 sex:女]]

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
	// 输出 =》 map[1:map[age:20 name:王武] 2:map[age:24 name:王六] 3:map[age:30 name:王七]]

	// 方式1
	var cat Cat
	fmt.Println(cat) // 输出 =》{ 0  <nil>}

	// 方式2
	cat2 := Cat{}     // 与 &Cat{} 写法有什么差别？
	fmt.Println(cat2) // 输出 =》{ 0  <nil>}

	// 方式3
	// new一个指针，引用类型
	cat3 := new(Cat) //等价 （var cat3 *Cat = new(Cat)）
	(*cat3).Name = "aaa"
	cat3.Name = "bbb"  // 两种写法都可以，底层转化为=》(*cat3).Name = "bbb"（注意：这是语法糖？它应该如何写？）
	fmt.Println(*cat3) // 输出 =》{bbb 0  <nil>}

	// 方式4
	cat4 := &Cat{}
	(*cat4).Name = "111"
	fmt.Println(*cat4) // =》{111 0  <nil>}

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
