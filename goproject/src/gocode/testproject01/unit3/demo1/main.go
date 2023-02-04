package main
import(
	"fmt"
)
func main(){
	//运算符
	var n1 int = +10
	fmt.Println(n1)//10
	var n2 int = 7 + 2
	fmt.Println(n2)//9
	var s1 string = "eef"+"ee23"
	fmt.Println(s1)//eefee23

	fmt.Println(10/3) //3//两个int类型数据运算，结果一定为整数类型
	fmt.Println(10.0/3)//3.3333333333333335	//浮点类型参与运算，结果为浮点类型

	//取模运算
	fmt.Println(10%3)//1
	fmt.Println(-10%3)//-1
	fmt.Println(10%-3)//1
	fmt.Println(-10%-3)//-1

	//自增操作
	var a int = 10
	a++
	fmt.Println(a) //11
	a--
	fmt.Println(a) //10
	//++ 自增 加1操作，--自减，减1操作
    //go语言里，++，--操作非常简单，只能单独使用，不能参与到运算中去
	//go语言里，++，--只能在变量的后面，不能写在变量的前面 --a  ++a  错误写法




	


}