package main
import(
	"fmt"
)
/* 
【1】切片(slice)是golang中一种特有的数据类型
【2】数组有特定的用处，但是却有一些呆板(数组长度固定不可变)，所以在 Go 语言的代码里并不是特别常见。相对的切片却是随处可见的，切片是一种建立在数组类型之上的抽象，它构建在数组之上并且提供更强大的能力和便捷。
【3】切片(slice)是对数组一个连续片段的引用，所以切片是一个引用类型。这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。切片提供了一个相关数组的动态窗口。
【4】代码：
切片的语法：
var 切片名 []类型 = 数组的一个片段引用
*/
func main() {
	//定义数组
	var intarr [6]int = [6]int{34,4,5,12,3,5}
	//切片构建在数组上
	//定义切片名slice []动态变化的数组长度不写int类型intarr是原数组
	//[1:3]切片 切出的片段从1开始到3结束不包括3 [1:3)
	//var slice []int = intarr[1:3]
	/* 
		简写方式：
		1) var slice = arr[0:end]  ----》 var slice = arr[:end]
		2) var slice = arr[start:len(arr)]  ----》  var slice = arr[start:]
		3) var slice = arr[0:len(arr)]   ----》 var slice = arr[:]
	*/
	slice := intarr[1:3]
	//输出数组
	fmt.Println("intarr",intarr)//intarr [34 4 5 12 3 5]
	//输出切片
	fmt.Println("slice",slice)//slice [4 5]

	//切片元素个数
	fmt.Println("slice元素个数:",len(slice))//slice元素个数: 2
	//切片容量，容量可以动态变化 大约是个数的二倍
	fmt.Println("slice的容量:",cap(slice))//slice的容量: 5

	//内存分析
	//切片有3个字段的数据结构：一个是指向底层数组的指针，一个是切片的长度，一个是切片的容量。
	fmt.Printf("数组下标为1的地址%p\n",&intarr[1])//数组下标为1的地址0xc00000e368
	fmt.Printf("切片下标为0的地址%p\n",&slice[0])//切片下标为0的地址0xc00000e368
	fmt.Printf("切片的地址%p\n",&slice)//切片的地址0xc000008078

	slice[1] = 55
	//切片改变了数组
	fmt.Println("intarr",intarr)//intarr [34 4 55 12 3 5]
	fmt.Println("slice",slice)//slice [4 55]
	//数组改变切片
	intarr[1] = 44
	fmt.Println("intarr",intarr)//intarr [34 44 55 12 3 5]
	fmt.Println("slice",slice)//slice [44 55]


}