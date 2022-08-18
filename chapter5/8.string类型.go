package chapter5

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 为什么 string 类型要单独拿出来分析：
// 因为 string 是引用类型，在使用上确跟值一样，这跟它的实现原理有关，且 string 在整个语言中极其重要。

// 基础定义
// 在Go语言中，字符串不同于其他语言，如Java、c++、Python等。
// 它是一个变宽字符序列，其中每个字符都用UTF-8编码的一个或多个字节表示。
// 或者换句话说，字符串是任意字节(包括值为零的字节)的不可变链，或者字符串是一个只读字节片，字符串的字节可以使用UTF-8编码在Unicode文本中表示。
// 由于采用UTF-8编码，Go字符串可以包含文本，文本是世界上任何语言的混合，而不会造成页面的混乱和限制。

// 字符串原理：
// string 在底层实现上是引用类型（两个字节，一个存对应字符类型的起始位地址，别一个指定其长度），但是因为 string 不允许修改，只能生成新的对象，在逻辑上和值类型无差别。
// 这种特性，直接导致 string 明明是引用类型，在使用上表现的跟值类型一样，这里有很多自动转换的机制支持。

func sss(ss1 string) {
	//ss1[1] = '5' // Go语言不允许修改 string 类型（只读），所以引用的特性体现不出来，永远是单向传递
	ss1 = "new string" // 又因为，在 Go 中的引用与其它语言 C++ 引用不同，在 Go 中，引用变量可以被整体覆盖
}

func Example5_8() {
	fmt.Println("Example5_8:")

	// 字符串是引用类型
	ss2 := "abcdedf"
	sss(ss2)
	fmt.Println(ss2)

	// 1.普通字符串，字符串字面量使用双引号（“”）创建，此类字符串支持转义字符
	s1 := "夏又桥"
	fmt.Println(s1)

	// 2.字符串文字是使用反引号（``）创建的，也称为raw literals(原始文本)。原始文本不支持转义字符，可以跨越多行，并且可以包含除反引号之外的任何字符。通常，它用于在正则表达式和HTML中编写多行消息。包裹能力极强
	s2 := `abc
ddd "abcd"`
	fmt.Println(s2)

	// 3.字符串不可变，一旦创建了字符串便不可变，无法更改字符串的值。换句话说，字符串是只读的。
	s3 := "夏又桥"
	for i, v := range s3 {
		fmt.Printf("%c 索引 -> %d\n", v, i) // 索引是存储UTF-8编码代码点的第一个字节的变量
	}
	//s3[0] = 'xia' // 无法修改
	for i, v := range "abc" {
		fmt.Printf("%c 索引 -> %d\n", v, i)
	}

	// 中文和字符读取对比
	fmt.Printf("%c\n", "abc"[0]) // a
	fmt.Printf("%c \n", s3[0])   // 乱码：å，因为一个文字由多个字节组成

	// 4.从切片创建字符串（强制转化）
	myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	fmt.Println("String 1: ", string(myslice1))

	myslice2 := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	fmt.Println("String 2: ", string(myslice2))

	// 5.字符串长度、字节数
	fmt.Println("abc 字节长：", len("abc"))
	fmt.Println("夏又桥 字节长：", len("夏又桥"))

	fmt.Println("abc 字符长：", utf8.RuneCountInString("abc"))
	fmt.Println("夏又桥 字符长：", utf8.RuneCountInString("夏又桥"))

	// 6.将切片字符串或数组字符串连接起来
	myslice := []string{"Welcome", "To", "nhooo", "Portal"}
	fmt.Println(strings.Join(myslice, "-"))

}

//17、字符串与[]byte之间的转换是复制（有内存损耗），可以用map[string] []byte建立字符串与[]byte之间映射，也可range来避免内存分配来提高性能
//for i,v := range []byte(str) {
//}

//23、使用for range迭代String，是以rune来迭代的。

//for range总是尝试将字符串解析成utf8的文本，对于它无法解析的字节，它会返回oxfffd的rune字符。
//因此，任何包含非utf8的文本，一定要先将其转换成字符切片([]byte)。
//data := "A\xfe\x02\xff\x04"
//for _,v := range data {
//fmt.Printf("%#x ",v)
//}
////prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)
//fmt.Println()
//for _,v := range []byte(data) {
//fmt.Printf("%#x ",v)
//}
//prints: 0x41 0xfe 0x2 0xff 0x4 (good)
