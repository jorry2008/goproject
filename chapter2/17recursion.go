package chapter2

import "fmt"

func Example17() {
	fmt.Println("Example:")

	// 顺序打印
	i := 5
	test1(i)

	// 阶乘
	var ii int = 30
	fmt.Printf("%d 的阶乘是:", ii)
	fmt.Printf("=%d \n", Factorial(uint64(ii)))
}

func test1(i int) {
	if i < 10 {
		i++
		println("当前i的值：", i)
		test1(i)
	}
}

// Factorial 分析递归时，只要看表达式本身就可以了
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		if n == 1 {
			fmt.Printf("%d", n)
		} else {
			fmt.Printf("%dx", n)
		}

		return n * Factorial(n-1) // 这个表达式表示：n会重复乘以它下面的n-1，直到n减小到0
	}

	return 1
}

// 递归用起来怎么那么像闭包
