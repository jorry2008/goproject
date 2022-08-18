package chapter5

// 常见的以下方式定义结构体的默认值（golang默认值无法在结构体中设计）

type Person struct {
	Name   string
	Age    int
	Weight int
	Foo    string
}

func NewPerson1() *Person {
	return &Person{Foo: "Person"}
}

// or

func NewPerson2() *Person {
	p := new(Person)
	p.Foo = "Person"
	return p
}

//or

func NewPerson3() Person {
	return Person{Foo: "Person"}
}
