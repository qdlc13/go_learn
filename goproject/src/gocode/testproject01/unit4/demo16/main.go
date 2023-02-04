package main
import(
	"fmt"
)
/* 什么是内置函数/内建函数：
Golang设计者为了编程方便，提供了一些函数，这些函数不用导包可以直接使用，我们称为Go的内置函数/内建函数。
内置函数存放位置：
在builtin包下，使用内置函数也的，直接用就行
常用函数：
*/
func main() {
	//len函数： 统计v的长度,按字节进行统计 v可以是数组切片映射等等
	//定义字符串
	str := "go lang 你好"
	fmt.Println(len(str))//14

	//new函数：func new(Type) *Type
	//内建函数new分配内存 其第一个实参为类型，而非值。 返回值是指向该类型的新分配的零值的指针
	//分配内存，主要用来分配值类型（int系列, float系列, bool, string、数组和结构体struct）

	num := new(int)   //num 是 *int
	fmt.Printf("num的类型是:%T,num的值是:%v,num的地址是:%v,num指向内存的值是:%v",num,num,&num,*num)
	//num的类型是L:*int,num的值是:0xc0000160a8,num的地址是:0xc00000a030,num指向内存的值是:0

	//make函数：
	//分配内存，主要用来分配引用类型（指针、slice切片、map、管道chan、interface 等）

}