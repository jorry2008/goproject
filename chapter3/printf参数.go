package chapter3

import (
	"fmt"
)

type S1 struct {
	Name string
}
type I1 interface {
	func1() string
}

type s1 struct {
	name string
}

func Example29() {
	fmt.Println("Example29:")

	// Go语言的标准输出流在打印到屏幕时有些参数跟别的语言（比如C#和Java）不同，具体如下：
	// Sprintf 格式化，占位符说明

	// 一、General

	//`%v` 以默认的方式打印变量的值（变量值）
	//`%T` 打印变量的类型（变量类型）
	t1 := 125
	s2 := S1{Name: "name"}
	fmt.Println("---------- 常规 ----------")
	fmt.Printf("默认的变量值：%v \n", t1)  // 125
	fmt.Printf("默认的变量类型：%T \n", t1) // int
	fmt.Printf("默认的变量值：%v \n", s2)  // {name}
	fmt.Printf("默认的变量类型：%T \n", s2) // chapter3.S1

	// 二、Integer

	//`%d` 打印整型
	//`%+d` 带符号的整型
	//`%o` 不带零的八进制
	//`%#o` 带零的八进制
	//`%q` 打印单引号
	//`%x` 小写的十六进制
	//`%X` 大写的十六进制
	//`%#x` 带 0x 的十六进制
	//`%U` 打印 Unicode 字符
	//`%#U` 打印带字符的 Unicode
	//`%b` 打印整型的二进制
	i1 := 20
	fmt.Println("---------- 整型 ----------")
	fmt.Printf("整型的值：%d \n", i1)           // 20
	fmt.Printf("整型的值：%d \n", i1*-1)        // 20
	fmt.Printf("带符号整型：%+d \n", i1)         // +20
	fmt.Printf("不带零八进制：%o \n", i1)         // 24
	fmt.Printf("带零八进制：%#o \n", i1)         // 024
	fmt.Printf("单引号：%q \n", i1)            // '\x14' ？？
	fmt.Printf("小写十六进制：%x \n", i1)         // 14
	fmt.Printf("大写十六进制：%X \n", i1)         // 14
	fmt.Printf("带0x的十六进制：%#x \n", i1)      // 0x14
	fmt.Printf("Unicode字符：%U \n", i1)      // U+0014
	fmt.Printf("带字符的Unicode字符：%#U \n", i1) // U+0014
	fmt.Printf("整型二进制：%b \n", i1)          // 10100

	// 三、Integer width

	//`%5d` 表示该整型最大长度是5，下面这段代码

	fmt.Println("---------- 整型等宽 ----------")
	fmt.Printf("|%5d| \n", 12)           // |   12|
	fmt.Printf("%5d \n", []int{12, 25})  // 默认右对齐 [   12    25]
	fmt.Printf("%-5d \n", []int{12, 25}) // -号左对齐 [12    25   ]

	// 输出结果如下：
	// |    1|
	// |1234567|

	//`%-5d`则相反，打印结果会自动左对齐
	//`%05d`会在数字前面补零

	// 四、Float

	//`%f` (=`%.6f`) 6位小数点
	//`%e` (=`%.6e`) 6位小数点（科学计数法）
	//`%g` 用最少的数字来表示
	//`%.3g` 最多3位**数字**来表示
	//`%.3f` 最多3位**小数**来表示
	fmt.Println("---------- 浮点 ----------")

	// 五、String

	//`%s` 正常输出字符串
	//`%q` 字符串带双引号，字符串中的引号带转义符
	//`%#q` 字符串带反引号，如果字符串内有反引号，就用双引号代替
	//`%x` 将字符串转换为小写的16进制格式
	//`%X` 将字符串转换为大写的16进制格式
	//`% x` 带空格的16进制格式
	fmt.Println("---------- 字符串 ----------")
	fmt.Printf("|%-5s| \n", "abc") // |abc  |

	// 六、String Width (以5做例子）

	//`%5s` 最小宽度为5
	//`%-5s` 最小宽度为5（左对齐）
	//`%.5s` 最大宽度为5
	//`%5.7s` 最小宽度为5，最大宽度为7
	//`%-5.7s` 最小宽度为5，最大宽度为7（左对齐）
	//`%5.3s` 如果宽度大于3，则截断
	//`%05s` 如果宽度小于5，就会在字符串前面补零
	fmt.Println("---------- 字符串等宽 ----------")

	// 七、Struct

	//`%v` 正常打印，比如：`{sam {12345 67890}}`
	//`%+v` 带字段名称，比如：`{name:sam phone:{mobile:12345 office:67890}`
	//`%#v` 用Go的语法打印，比如`main.People{name:"sam", phone:main.Phone{mobile:"12345", office:"67890"}}`
	fmt.Println("---------- 结构体 ----------")

	// 八、Boolean

	//`%t` 打印 true 或 false
	fmt.Println("---------- 布尔 ----------")

	// 九、Pointer

	//`%p` 带 0x 的指针
	//`%#p` 不带 0x 的指针
	fmt.Println("---------- 指针 ----------")

	fmt.Printf("%d \n", 10)              // 正常输出 10
	fmt.Printf("%#o \n", 10)             // 八进制加0 012
	fmt.Printf("%#x \n", 10)             // 十六进制加0x 0xa
	fmt.Printf("%#X \n", 10)             // 十六进制加0X 0XA
	fmt.Printf("%#p \n", &S1{Name: "a"}) // c00003a270
	fmt.Printf("%p \n", &S1{Name: "a"})  // 0xc00003a280 // #表示去掉指针地址的 0x

	fmt.Printf("% d", []int{21, 25, 14}) // [ 21  25  14]

}
