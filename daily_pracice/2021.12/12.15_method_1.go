package main

import "fmt"

type ByteSlice []byte

func main() {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
}

//This rule arises because pointer methods can modify the receiver; invoking them on a value
//would cause the method to receive a copy of the value, so any modifications would be
//discarded. The language therefore disallows this mistake. There is a handy exception,
//though. When the value is addressable, the language takes care of the common case of
//invoking a pointer method on a value by inserting the address operator automatically. In our
//example, the variable b is addressable, so we can call its Write method with just b.Write.
//The compiler will rewrite that to (&b).Write for us.
//之所以会有这条规则是因为指针方法可以修改接收者；通过值调用它们会导致方法接收到该
//值的副本， 因此任何修改都将被丢弃，因此该语言不允许这种错误。不过有个方便的例外：
//若该值是可寻址的， 那么该语言就会自动插入取址操作符来对付一般的通过值调用的指针方
//法。在我们的例子中，变量 b 是可寻址的，因此我们只需通过 b.Write 来调用它的 Write 方
//法，编译器会将它重写为 (&b).Write。
func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	slice = append(slice, data...)
	*p = slice
	return len(data), nil
}
