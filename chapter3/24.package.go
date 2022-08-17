package chapter3

// 1.包是 go 中项目依赖或文件管理的最小单元（其它 oop 语言，所有元素都在类中，所有相关类都在一个包中，包依赖在vender代码仓下）
// 2.在包的规则下，只有文件夹和包名是关键点，文件名似乎被忽略了？
// 2.一个包就是一个文件夹，文件夹可以有多个go文件或多个子文件夹，通常包名与包文件夹一致，也可以不同，只是在引入操作上会有一些差异（不同的话，容易导致使用方误解，极不方便）
// 3.
import (
	"fmt" // 包名与文件夹名相同时：项目名/子文件1/子文件2....
	"github.com/jorry2008/goproject-demo/pk1"
	//pack2 "test/test2" // 包名与子目录不同时：包名 项目名/子文件1/子文件2....
)

func Example23() {
	fmt.Println("Example23:")

	dog := pk1.Dog{Name: "jorry", Age: 9, Other: map[string]string{
		"a": "a",
		"b": "b",
	}}

	fmt.Println(dog)
}
