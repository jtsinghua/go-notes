package main

import "fmt"

//指针
func main() {
	var num int = 0
	//指针声明与初始化
	//默认为nil
	//操作符 "&" 取变量地址， "*" 透过指针访问⺫标对象
	var t *int = &num

	fmt.Println(t) //0xc000060080

	//不⽀持指针运算，不⽀持 "->" 运算符，直接⽤ "." 访问⺫标成员
	type data struct {
		age int
	}

	var dataVar *data = &data{
		age: 10,
	}
	//不⽀持指针运算，不⽀持 "->" 运算符，直接⽤ "." 访问⺫标成员
	fmt.Println(dataVar.age)
	//不能对指针做加减法等运算
	str := "this"
	var ptr *string = &str
	//ptr++
	fmt.Printf("ptr = %p\n", ptr)
}
