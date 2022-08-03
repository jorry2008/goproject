package chapter2

import "fmt"

// Map 是一种无序的键值对的集合
// Map 只能通过 key 来快速检索数值
// Map 的实现原理是基于 hash 表，所以效率高，但无法保证读取顺序！
// delete() 是集合专用删除内置函数

// 声明变量，默认 map 是 nil
// 如果不初始化 map，那么就会创建一个 nil map，而 nil map 不能用来存放键值对

var varmap1 map[string]string                                                         // 只声明，不初始化为 nil，没有分配空间
var varmap2 = map[string]string{"aa": "aa", "bb": "bb", "cc": "cc"}                   // 声明并初始化
var varmap3 map[string]string = map[string]string{"aa": "aa", "bb": "bb", "cc": "cc"} // 完整写法

// 推荐写法
var varmap4 = map[string]string{} // 初始化为 {} 有效，可用

func Example13() {
	println("Example13:")

	varmap4["aa"] = "aa" // 声明，并初始化为 {} 空，可使用
	// 或
	varmap5 := make(map[string]int)
	println("使用make创建的集合，正常分配内存：", &varmap5)
	println("make创建的集合是否为nil：", varmap5 == nil) // 返回 false
	varmap5["aa"] = 11                          // 分配了内存，且可用

	println(varmap1)            // 0x0，没有分配内存空间
	fmt.Println(varmap1 == nil) // 整个集合就是 nil，无法直接存储键值对！！！

	var countryCapitalMap map[string]string // 决定集合类型的包括三部分：map关键字，string键，string值
	//countryCapitalMap["aabc"] = "ddddd"     // idea 没报错，但编译不通过
	countryCapitalMap = make(map[string]string) // 只声明的 nil 集合，可以接受其它同类型的集合！
	// 或
	// countryCapitalMap := make(map[string]string)

	countryCapitalMap["France"] = "巴黎" // 原声明的集合，无法使用键值对，但可以接受其它同类型集合赋值，赋值后，仍然可以使用键值对！！！
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["China"] = "中国"

	// 以下只返回一个键
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 以下特性，集合专用。。。
	capital, ok := countryCapitalMap["American"] // 请集合时，如果不存在，会返回别一个参数表示是否存在
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println(ok) // 这是一个布尔值（牛逼啊。。。）
		fmt.Println("American 的首都不存在")
	}

	// 删除操作
	delete(countryCapitalMap, "China") // 通过 key，删除一个集合中的元素
	capital2, ok2 := countryCapitalMap["China"]
	if ok2 {
		println("未删除", capital2)
	} else {
		println("China 已经删除")
	}
}

// 小结：
// 1.只声明的 map 没有分配内存其值为 nil，且无法直接使用，但能接受其它 map 数据
// 2.声明并初始化为 {}，会分配内存，可正常使用
// 3.使用make()函数，会分配内存，可正常使用
// 4.只有集合取元素时，如果不存在，同时额外返回一个布尔值作为第二返回值
