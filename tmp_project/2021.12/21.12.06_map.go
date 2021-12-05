package main

import (
	"fmt"
)

func main() {
	// 创建了一个预估容量是2的 map
	m := make(map[string]string, 2)
	m["hello"] = "world"
	m["hello1"] = "world"
	m["hello2"] = "world" // 超过容量不抛错
	for k, v := range m {
		fmt.Printf("m: %s = %s \n", k, v)
	}

	// 没有指定容量
	m1 := make(map[string]string)
	m1["hello"] = "world"
	for k, v := range m1 {
		fmt.Printf("m1: %s = %s \n", k, v)
	}

	_, ok := m["不存在的key"]
	if !ok {
		println("key不存在")
	}

	// 直接初始化
	m2 := map[string]string{
		"Tom": "Jerry",
	}
	for k, v := range m2 {
		fmt.Printf("m2: %s => %s \n", k, v)
	}

}
