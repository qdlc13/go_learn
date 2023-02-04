package main
import(
	"fmt"
)

/* 
	为了简化数据类型定义,Go支持自定义数据类型
	基本语法: type 自定义数据类型名  数据类型       
	可以理解为 : 相当于起了一个别名
	例如:type mylnt int -----》这时mylnt就等价int来使用了.
*/
func main() {
	//自定义数据类型相当于取别名
	type myInt int

	var num1 myInt = 20
	fmt.Println("num1:",num1) //num1: 20
	var num2 int = 10
	/* num2 = num1 报错 虽然是取别名但是go中编译识别认为myInt和int不是一种数据类型
		# command-line-arguments
		.\main.go:17:9: cannot use num1 (variable of type mylnt) 
		as type int in assignment
	*/
	num2 = int (num1)
	fmt.Println("num2:",num2)//num2: 20
}