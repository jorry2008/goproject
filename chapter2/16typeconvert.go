package chapter2

import "fmt"

// 类型转换：自动转换，强制转换，精度问题
// 整型、布尔、浮点、字符串、虚数 基本类型的转换

func Example16() {
	println("Example16:")

	// 自动转换
	fmt.Printf("%T \n", 25.2+10) // float64，自动转化，保住了精度

	// 强制转换
	// 类型转换用于将一种数据类型的变量转换为另外一种类型的变量，
	// type_name(expression)
	// 布尔类型，与整型之间不能转化

	var a int64 = 3
	var b int32
	b = int32(a) // 强制转换为 int32，才能进行同类型赋值操作
	fmt.Printf("b 为 : %d \n", b)

	varint := 25
	varfloat := 25.2
	varstring := "测试"

	println("整型25转为浮点：", float32(varint))
	println("浮点转化为整型 ：", int(varfloat)) // 会有精度损失

	for sk, sv := range []byte(varstring) {
		println(sk, sv)
	}
	bytess := []byte(varstring)
	fmt.Printf("%q", bytess[:3]) // 返回：测
}
