package chapter2

import "fmt"

func Example15() {

	printStr("hello")
	fmt.Println()
	fmt.Println()
	printStr("中国人")
}

func printStr(s string) {
	fmt.Println("str: " + s)
	for _, v := range s {
		fmt.Printf("0x%x %c, ", v, v)
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x, ", s[i])
	}
}
