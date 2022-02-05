package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			// i是共享的
			// 方法中的参数都是值传递，func(i)中的i被复制一份
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond * 50)
}
