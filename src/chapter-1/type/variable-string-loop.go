package main

import "fmt"

func main() {
	str := "this is a test"
	//两种遍历
	//byte
	for i := 0; i < len(str); i++ {
		fmt.Print(str[i], " ")
	}
	fmt.Println()

	//rune
	for _, n := range str {
		//fmt.Printf("%T\n", n) //int32
		fmt.Printf("%c ", n)
	}
	fmt.Println()
}
