package chapter5

import "fmt"

// 将 slice、map、struct 三个引用类型，混合使用

type stu struct {
	id   int
	name string
	age  int
}

func SliceMapStruct() {
	fmt.Println("SliceMapStruct:")

	// 一、map 嵌套
	students := make(map[int]map[string]string)
	// 赋值方式1
	students[0] = make(map[string]string) // 形式1
	students[1] = map[string]string{}     // 形式2
	// 赋值方式2
	students1 := map[int]map[string]string{
		0: {"name": "张三",
			"sex": "男"},
		1: {"name": "张四",
			"sex": "男"},
	}
	fmt.Println(students)
	fmt.Println(students1)

	// 二、slice + map
	smap := make([]map[string]int, 2, 4) // 定义一个切片，元素为map集合
	// 赋值方式1
	smap[0] = make(map[string]int)
	smap[0]["a"] = 1
	// 赋值方式2
	smap[1] = map[string]int{"a": 0, "b": 1}
	// 切片追加
	smap = append(smap, map[string]int{"c": 2})
	fmt.Println(smap)

	// 三、map + struct
	mapstruct := map[int]stu{}

	stu1 := stu{
		id:   1,
		name: "张三",
		age:  20,
	}
	stu2 := stu{
		id:   2,
		name: "张三",
		age:  21,
	}
	mapstruct[0] = stu1
	mapstruct[1] = stu2
	fmt.Println(mapstruct)
	fmt.Println(mapstruct[0].name)

	// 四、struct + slice
	structslice := []stu{0: {
		id:   1,
		name: "王五",
		age:  20,
	}}
	// 切片追加
	structslice = append(structslice, stu1)
	fmt.Println(structslice)
}
