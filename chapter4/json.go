package chapter4

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name string `json:"name"` //标记json名字为name
	Age  int    `json:"age"`
	Time int64  `json:"-"` // 标记忽略该字段
}

func Test55() {
	person := Person{"小明", 18, time.Now().Unix()}
	if result, err := json.Marshal(&person); err == nil {
		fmt.Println(string(result))
	}
}
