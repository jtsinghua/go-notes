package main

import "fmt"

type Demo1 struct {
	name string
	age  int
}

func main() {
	//定义变量关键字
	var iNum int
	var fNum float32
	var fNum64 float64
	//如果定义变量的时候没有初始化，则自动初始化为零值
	//这里，零值对于不同数据类型，各有不同
	//对于int类型，为0
	fmt.Printf("iNum = %d\n", iNum) //iNum = 0
	//对于float,则为0.000000
	fmt.Printf("fNum = %f\n", fNum)     //fNum = 0.000000
	fmt.Printf("fNum64 = %f\n", fNum64) //fNum64 = 0.000000

	var str string
	//字符串自动初始化为""
	fmt.Printf("str = %s\n", str) //str = ""
	fmt.Println(str == "")        //true

	var d Demo1
	//类型的自动初始化为各属性的自动初始化
	fmt.Printf("d = %v\n", d) //d = { 0}
}
