package main

import "fmt"

//常量
//变量必须在编译期就可确定值
//定义常量组
const (
	HOST       = "localhost"
	USERNAME   = "scott"
	isReadable = false
	TIMEOUT    = 1000
)

//在常量组中，若是没有指定类型或值，则与上一个相同
//这就是说salary1是int类型，且值为10000
const (
	salary0 = 10000
	salary1
	salary2 = 18000
)

func main() {
	fmt.Printf("host = %s, username = %s, isReadable = %v, timeout = %d\n", HOST, USERNAME, isReadable, TIMEOUT)
	//常量若是重新赋值，则编译器报错
	//isReadable = true //cannot assign to isReadable
}
