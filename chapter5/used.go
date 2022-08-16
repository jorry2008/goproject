package chapter5

// 来自 Go 语言的怒吼：声明了一堆、初始化了一堆，你不使用是干啥玩意儿？

func Example00007() {
	println("Example00007:")

	// 一个变量，声明了不算被使用了，局部变量会报错
	// 一个变量，声明了，且初始化了，也不算使用了，局部变量仍然会报错
	var ptr1 *int
	var pptr2 **int // 到这一步，ptr1 和 pptr2 只量声明了，没有被使用到
	ppvalue := 25   // 到这一步，ppvalue 声明且初始化了，但仍然是没有被使用

	ptr1 = &ppvalue // 到这一步，ptr1 初始化了，而 ppvalue 算是使用了，因此 ptr1 还是会报错
	pptr2 = &ptr1   // 到这一步，pptr2 初始化了，而 ptr1 算是使用了，因此 pptr2 还是会报错

	println("指向指针的指针地址：", pptr2) // 到这一步，才算所有变量都被使用了！！！

}
