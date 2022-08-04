package chapter2

import (
	"errors"
	"fmt"
	"runtime/debug"
)

// go 的错误（异常）处理比较特别：
// 默认情况下，在 Go 中出现异常不处理也不会影响代码流程，需要结合 panic 实现程序终止
// 异常是一个接口，可以在 Go 中自定义错误类型，通过自定义一个结构体进而实现 error 接口来增加我们的错误信息（通过接口，我们可以实现任何想要类型的错误返回）
// go 原生包装了 errors.New("") 方法，本质还是 error 接口的实现

/*
go 内置接口如下：
type error interface { // error 本身就是接口类型
	Error() string
}

原生包含的 errors.New() 方法：
type errorString struct { // 超级简单的结构体
	s string
}
func (e *errorString) Error() string {
	return e.s
}
func New(text string) error {
	return &errorString{text}
}
*/

// 使用时注意事项:
// 1.在返回值中包含有 error 的情况，error 应该放在最后返回
// 2.recover() 函数能收集到 panic() 发出来的异常信息并进行处理（类似于 try catch，只是更强）
// 3.defer语句，总是会在 fn 函数结束之后才执行（设置延迟执行）
// 4.defer 需要再 panic 之前声明，否则由于 panic 之后的代码得不到执行
// 5.发生 panic 后，程序会从调用 panic 的函数位置或发生 panic 的地方立即返回，逐层向上执行函数的 defer 语句，然后逐层打印函数调用堆栈，直到被 recover 捕获或运行到最外层函数
// 6.panic 不但可以在函数正常流程中抛出，在 defer 逻辑里也可以再次调用 panic 或抛出 panic
// 7.defer 里面的 panic 能够被后续执行的 defer 捕获
// 8.recover 用来捕获 panic，阻止 panic 继续向上传递
// 9.recover()和defer一起使用，但是 defer 只有在后面的函数体内直接被掉用才能捕获 panic 来终止异常，否则返回nil，异常继续向外传递
// 10.对于一个应用而言，错误的结构应该是统一的，而非随便定义

// 小节：
// panic() 一旦被抛出，会一层一层向上传递的，直到捕获或导致程序最终失败；
// recover() 能接受到被 defer 定义之后的 panic 异常；
// defer 定义是逐级由内向外一层一层生效的；

// 比如，可以在一个函数中设置 defer 匿名函数，函数中使用 recover 接受异常，接着在其后的逻辑中抛出 panic，这样就可以调试这个函数，其外层程序仍然可以继续执行！
// 再比如，如果局部有致命错误，则在局部的 defer 中还可以继续 panic 向外抛出，直到整个程序失败，或者将异常抛出到最外层的 recover 进行统一接受并处理；

func Example19() {
	println("Example19:")

	// 方式一：直接返回错误，使用原生包装方法（官方提供的简易字符串结构体）
	err := TryErrorFn()
	fmt.Println(err) //

	// 方式二：实现 error 接口（自定义错误信息结构体）
	_, e := TryErr(1)
	fmt.Println(e) // 在其它语言中，抛出异常就终止了，但在 go 中，不处理就不终止

	// 错误与终止
	// func panic(interface{}) // panic() 函数接受一个空类型，即接受一切类型，程序执行到 panic 会终止执行，并抛出异常
	// func recover() interface{} // recover() 同上，函数来捕获 panic() 函数的异常
	// defer语句的作用就是延迟调用函数，等到函数返回的时候在调用，一个函数中可以包含多个，执行顺序是后入为主的顺序执行
	// panic -> recover 如同一个错误通道，然后 defer 将代码块设置到程序在任何情况下执行完成都会调用，三者之间就构建出一个专用的异常信息流通道！

	// 以下匿名函数，会在整个程序运行完成后被调用【defer和recover结合，可以设计出全局/或局部异常捕获的能力，还能依据不同类型选择处理的方式】
	//defer func() {
	//	if err := recover(); err != nil { // 主要做的是：指定作用域下错误收集，并自定义处理
	//		fmt.Println("捕获 Panic 抛出的错误：", err)
	//	}
	//	time.Sleep(3 * time.Second) // 延时执行
	//	fmt.Println("program done....")
	//}()

	// 正常的业务代码，抛出各种 panic
	//panic("panic1")             // 在任何时候，只要 panic 抛出，程序就会终止，但终止前，它会主动调用 defer 匿名函数，并执行一遍！
	//fmt.Println("after panic1") // 不会执行（idea也将其标记了无效代码）

	// 设计一个完整的异常构架
	// 1.一个全局错误结构体：code、type、message
	// 2.定义统一的错误显示格式
	// 3.统一的异常处理机制
	defer func() {
		if errStruct := recover(); errStruct != nil {
			switch errStruct.(type) { // 断言写法
			case ErrorWarning:
				if err2, ok := errStruct.(ErrorWarning); ok { // 断言类型转换
					println(err2.Error())
					debug.PrintStack()
				}
			case ErrorDanger:
				// 将错误转换为具体的类型（强制类型转换）
				if err2, ok := errStruct.(ErrorDanger); ok {
					println(err2.Error())
					debug.PrintStack()
				}
			default:
				fmt.Printf("%T \n", err)
			}
		}
	}()

	//panic(ErrorWarning{code: 419, message: "权限问题"})
	panic(ErrorDanger{code: 500, message: "程序运行问题"})

}

func TryErrorFn() error {
	return errors.New("这里有个错误（使用errors.New()）") // 不需要定义错误信息格式
}

type MyError struct { // 自定义的异常结构体，用于展示错误详情，并实现 error 接口
	code    int32
	message string
}

// 实现 error 接口的 Error 方法
func (e *MyError) Error() string { // 传入参数，即可以传值，也可以传引用？？？
	return fmt.Errorf("此程序出现了异常： code=%d, message=%s", e.code, e.message).Error()
}

func TryErr(a int64) (int64, error) {
	if a == 1 {
		return 1, &MyError{-1, "传入的参数错误"} // 因为返回类型为 error，所以这里返回的必须是实现接口的结构体的指针 &MyError
	}
	return 0, nil
}

type ErrorWarning struct {
	code    int
	message string
}

type ErrorDanger struct {
	code    int
	message string
}

func (error *ErrorWarning) Error() string {
	return fmt.Errorf("warngin异常： code=%d, message=%s", error.code, error.message).Error()
}

func (error *ErrorDanger) Error() string {
	return fmt.Errorf("danger异常： code=%d, message=%s", error.code, error.message).Error()
}

// 想要充分理解异常机制，看一下成熟的框架产品就可以了：局部调试、全局捕获管理
