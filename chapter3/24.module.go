package chapter3

// module github.com/jorry2008/goproject-demo
import "fmt"

// 1.模块：modules 是一个或多个放在一起的包的集合
// - 由于 go modules 不再依赖 GOPATH，所以项目可以放在任何地方——这个 『任何』，是指可以不是 GOPATH ，但是要是习惯了，继续放在 %GOPATH%/src 也没问题。
// - 进入 GOPATH/abc 下 使用 go mod init xxx 可创建任意名称的 module，文件夹名与模块名可以不一致
// 在 go 的项目依赖和管理中，最小管理单元是包，一个包有一个或多个go文件（不同于 php、java 语言，它们的最小管理单元是类）
// 在包模式下 包名是从GOPATH/src/ 后开始计算的，使用/ 进行路径分隔。
// 包管理有 Go Module 模式（不再由 GOPATH 限制），GOPATH 模式，对比两种模式

// 2.建议：

// 3.小结：

// https://morven.life/posts/golang-module/

// 如何写一个给其它人使用的模块？

func Example24() {
	fmt.Println("Example24:")

}
