package chapter2

import "fmt"

// 字符串和字节是可以相互转换的

func Example15() {
	fmt.Println("Example15:")

	// string to []byte
	s1 := "字符串和字节相互转换"
	b := []byte(s1) // 转换时只能用切片【三个字节，一个字符】
	// []byte to string
	s2 := string(b)
	println("切片大小：", len(s1), s1, s2)

	printStr("hello")
	fmt.Println()
	printStr("中国人")

	fmt.Println()
	str := "Go爱好者"
	fmt.Printf("The string: %q \n", str)               // %q 表示，显示字符串
	fmt.Printf("  => runes(char): %q \n", []rune(str)) // rune 是 int32，代表一个 Unicode 字符
	fmt.Printf("  => runes(hex): %x \n", []rune(str))  // %x 表示，每个字节用两字符十六进制数表示
	fmt.Printf("  => bytes(hex): [% x] \n", []byte(str))
}

func printStr(s string) {
	fmt.Println("str: " + s)
	for _, v := range s {
		fmt.Printf("0x%x -> %c, ", v, v)
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x -> %U (%T), ", s[i], s[i], s[i])
	}
}
