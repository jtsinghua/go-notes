package main

import "fmt"

//定义全局变量
var iAge int = 11

//重新赋值与定义同名的新变量
func main() {
	//局部变量，和全局变量IAge虽然同名，但不是同一个变量
	iAge := 12
	useGlobalVariable()             //全局变量iAge = 11
	fmt.Printf("iAge = %d\n", iAge) //iAge = 12

	{
		//这里又定义了一个新的同名变量
		iAge := 13
		fmt.Printf("iAge = %d\n", iAge)
	}
	fmt.Printf("iAge = %d\n", iAge)

}

//函数，使用全局变量
func useGlobalVariable() {
	fmt.Printf("全局变量iAge = %d\n", iAge)
}
