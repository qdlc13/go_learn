package main
import(
	"fmt"
)
/* 	【1】Go在不同类型的变量之间赋值时需要显式转换，并且只有显式转换(强制转换)。
	【2】语法：
	表达式T(v)将值v转换为类型T
	T : 就是数据类型
	v : 就是需要转换的变量
 */
func main(){
	//数据类型转换
	var n1 float32 = 100.002
	/* 
		var n2 int = n1  在这里自动转换不好使
		.\main.go:8:15: cannot use n1 (variable of type float32) 
		as type int in variable declaration
	*/
	
	fmt.Println(n1)  //100.002
	//转换
	var n2 int = int(n1)
	fmt.Println(n2) //100丢精度了 
	//n1类型不变
	fmt.Printf("%T\n",n1)//float32

	//int64 转 int8 编译不会出错可能会数据溢出
	var n3 int64 = 888888
    var n4 int8 = int8(n3)
    fmt.Println(n4)//56 溢出

	var n5 int32 = 12
	var n6 int64 = int64(n5) + 30 //一定要匹配=左右的数据类型 n5+30 报错
	fmt.Println(n5) //12
    fmt.Println(n6) //42

	//int8 范围-128~127
	var n7 int64 = 12
	var n8 int8 = int8(n7) + 127
	fmt.Println(n8)	//-117溢出	//编译通过，但是结果可能会溢出
	//128 (untyped int constant) overflows int8
	//var n9 int8 = int8(n7) + 128 //编译不会通过 因为128已经超过
	//fmt.Println(n9)

	var a int16 = 200
	var b int8 = int8(a)//编译能通过但是溢出
	fmt.Println(b) //-56

}