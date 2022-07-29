package chapter1

// 在包的规则下，只有文件夹和包名是关键点，文件名似乎被忽略了
import (
	"fmt" // 包名与文件夹名相同时：项目名/子文件1/子文件2....
	//pack2 "test/test2" // 包名与子目录不同时：包名 项目名/子文件1/子文件2....
)

func Example005() {
	fmt.Println()
}
