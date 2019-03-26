package main

import "fmt"

//变量
func main() {
	//如果提供了初始值，可以省略var关键字和类型，由编译器自动推断变量的类型
	//具体语法是 <变量名> := <初始值>
	iNum := 10
	fmt.Printf("iNum = %d\n", iNum)
	//可以一次定义多个变量
	//注意，其类型可以不同
	x, y, z := 1, "jerry", 0.01
	fmt.Printf("x = %d, y = %s, z = %f\n", x, y, z)
	//一次定义多个变量
	var (
		//自动初始化为0
		age int
		//自动初始化为空字符串
		name string
	)
	fmt.Printf("age = %d, name = %s\n", age, name)
	//对于多变量的赋值，先会从左到右计算右值,再从左到右赋值
	iAge, sName := getAge(), getName()
	fmt.Printf("iAge = %d, iName = %s\n", iAge, sName)

	//若是定义了变量，而没有使用，则编译时会报错
	//但只是对局部变量来说，对全局变量则编译器不会报错
	var demo string //demo declared and not used
	//可以用这种形式，规避编译器的检查
	_ = demo
	//当然，最常用的还是忽略函数的返回值
	//_，是一个特殊的变量，用于忽略值占位
	str, _ := returnTwoValue()
	fmt.Println(str)
}

//函数，返回两个值
//在go中，可以返回多个值
func returnTwoValue() (string, error) {
	return "", nil
}

//函数，返回一个字符串
func getName() string {
	fmt.Println("in getName()")
	return "jerry"
}

//函数，返回一个int值
func getAge() int {
	fmt.Println("in getAge()")
	return 10
}
