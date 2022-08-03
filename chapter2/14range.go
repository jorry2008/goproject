package chapter2

import (
	"fmt"
)

// Go 语言中 range 关键字用于 for 循环中迭代 数组(array)、切片(slice)、通道(channel)、集合(map) 的元素
// 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 键值对
// range 遍历返回的永远是两个值（多返回值），但两种接受方式

func Example14() {

	// 注意：数组、切片、集合、通道 都支持以下三种模式遍历
	//for key, value := range oldMap { // 返回 索引+值、键值
	//	// oldMap[key] == value
	//}
	//for key := range oldMap { // 返回 键、索引【为什么有这种方式，因为返回索引或键，可以进一步获取值，但只返回值是反向获取不到索引或键的】
	//	// oldMap[key]
	//}
	//for _, value := range oldMap { // 返回 值，将键或索引屏蔽掉
	//	// nothing
	//}

	// 以集合举例：

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
}
