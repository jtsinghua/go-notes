package main

//go中没有枚举类型，由常量来模拟
//iota关键字从0开始按行计数
const (
	Sunday    = iota //0
	Monday           //1
	Tuesday          //2
	Wednesday        //3
	Thursday         //4
	Friday           //5
	Saturday         //6
)

//可以定义多个iota，各自增长
const (
	level0, level1 = iota, iota << 10
	level2, level3
)

//若是iota被打断，则需要显式恢复
const (
	id0 = iota // 0
	id1
	id2 = "id2"
	id3 = iota //显式恢复 3
	id4        // 4
)

func main() {

}
