package main

import "fmt"

// string is the set of all strings of 8-bit bytes,
// conventionally but not necessarily representing UTF-8-encoded text. A string may be empty, but
// not nil. Values of string type are immutable.
//字符串是不可变类型
//内部用指针指向字节数组
func main() {
	//默认值是空字符串
	var str string
	fmt.Println(str)

	name := "jerry"
	//可以用索引值来访问某个字节
	fmt.Println(name[0]) //106
	//不能用&获取某个字节的指针
	//ptr := &name[0]
	//不支持修改
	//name[i] = 90
	//支持换行，不会转义
	str = `
this is a \n \t \b test
`
	fmt.Println(str)
	//连接跨⾏字符串时， "+" 必须在上⼀⾏末尾，否则导致编译错误
	str = "this is a " +
		"test"
	fmt.Println(str)
	//⽀持⽤两个索引号返回⼦串。⼦串依然指向原字节数组，仅修改了指针和⻓度属性
	subStr := str[0:4]
	fmt.Println(subStr)                                  //this
	fmt.Printf("str = %v, subStr = %v\n", &str, &subStr) //str = 0xc00004c1c0, subStr = 0xc00004c1f0

	//字符串不可变，但要改变也有办法
	bytes := []byte(str)
	bytes[0] = 'T'

	str = string(bytes)
	fmt.Println(str) //This is a test

	str = "字符串"
	runes := []rune(str)
	runes[2] = '牛'

	str = string(runes)
	fmt.Println(str) //字符牛
}
