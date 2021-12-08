package main

type animal interface {
	// Eat 这里可以有任意多个方法，不过我们一般建议是小接口
	//即接口里面不会有很多方法
	//方法声明不需要 func 关键字
	Eat()
}

type Duck interface {
	Swim()
}
