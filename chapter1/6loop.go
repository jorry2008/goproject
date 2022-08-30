package chapter1

import (
	"fmt"
)

func Example6() {
	fmt.Println("Example6:")

	// for 循环可以有三种形式，且不直接支持 while：
	// for init; condition; post {} // 标准的 c 语言形式
	// for condition {} // 同 while 用法一致
	// for {} // 无限循环

	// 循环控制语句
	// break	用于中断当前 for 循环或跳出 switch 语句
	// continue	跳过当前循环的剩余语句，然后继续进行下一轮循环
	// goto	    无条件跳转（预先将指定的语句块 statement 设置为自定义 label，然后由 goto 无条件跳转至语句块重复执行）

	// 标准用法
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("标准用法 0 ~ 10 之间的和：%d \n", sum)

	// 类 while 用法
	sum1 := 1
	for sum1 <= 10 {
		sum1 += sum1
	}
	fmt.Printf("类 while用法 1+1+2+4+8 之和：%d \n", sum1)

	// 无限循环
	//sum2 := 1
	//for {
	//	sum2++
	//}
	//fmt.Println(sum2)

	// 遍历操作
	// 所有的切片、map或者其它关联性数据结构，都使用 range 结构对其进行遍历操作
	strings := []string{"aaa", "bbb"} // range 遍历索引数组
	for k, v := range strings {
		println(k, v)
	}
	fmt.Println()
	numbers := [6]int{8, 9, 41, 7} // 注意，这个数组申请的空间大小是6，但实际存储的4个业务数据，最后两个是默认值占位0，实际仍然是6位数组
	for i, number := range numbers {
		println(i, number) // 遍历6次
	}

	// goto 操作
	a := 10
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
