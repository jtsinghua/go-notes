package main

import "fmt"

//引用类型
func main() {
	//数组
	array := []int{1, 2, 3}
	fmt.Println("array =", len(array))
	//切片
	sli := make([]int, 3)
	sli[0] = 1
	sli[1] = 2
	sli[2] = 3
	fmt.Println(sli)
	//map
	keyValue := make(map[string]string)
	keyValue["name"] = "jerry"
	keyValue["age"] = "10000"
	fmt.Println(keyValue)
}
