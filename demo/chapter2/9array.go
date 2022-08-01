package chapter2

import "fmt"

// Go 语言中有数组，只是没有 php 那么灵活，相比 Java 其实是差不多的
// 数组是具有"相同类型"的一组已编号且长度固定的数据项序列（连续内存空间），这种类型可以是任意的原始类型，例如：整型、字符串或者自定义类型，对数组整体而言类型为 [size]type,对数组元素而言每个元素都是原基本类型
// 重点：数组的定义是 [SIZE]variable_type，说明使用数组表达一个具体的类型的方式就是明确的数组大小和唯一的元素类型，两者关联在一起，才算是完整的类型！

//var numbers5 [5]int = [5]int{20} // 因为元素的类型为 int 初始化了一个 20，剩下了4个空位就当使用默认值 0 来填充【左边的5和右边的5必须保持一致，才被认为是同一个数据类型】
var numbers5 = [5]int{20}                 // 同上，数组明确的大小和元素类型，那么数组类型就明确了，所以变量后面不需要重复带类型
var balance1 [10]float32                  // 只声明不初始化，所有元素填充默认值
var balance2 = [5]float32{1: 2.0, 3: 7.0} // 选择性初始化，将索引为 1 和 3 的元素初始化，其它给默认值

//var varFloatArr = [...]float32{25, 36, 52.2}
var varFloatArr = []float32{25, 36, 52.2} // 同上，当数组大小不确认时，就不用填充 SIZE，不填充不代表没有，在编译器能力内会自动添加 SIZE【只是个快捷技巧】

// 注意：初始化数组中 {} 中的元素个数不能大于 [] 中的数字

func Example9() {
	println("Example9:")

	// 数组遍历
	balance2[4] = 25
	for i, f := range balance2 {
		println(i, f)
	}

	for k := 0; k < len(varFloatArr); k++ {
		fmt.Printf("varFloatArr[%d] = %f\n", k, varFloatArr[k])
	}

	// 多维数组

	// 数组作为函数参数

	// 精度问题

}
