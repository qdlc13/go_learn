 package main
 import(
	"fmt"
 )
 //补充：学习go的内存结构分析 栈 堆 代码区
 func exchangeNum(num1 int , num2 int) {
	var t int 
	t = num1
	num1 = num2
	num2 = t
 }
 func main(){
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("交换前的两个数num1 = %v,num2 = %v\n",num1,num2)
	//交换前的两个数num1 = 10,num2 = 20
	exchangeNum(num1,num2)
	fmt.Printf("交换后的两个数num1 = %v,num2 = %v\n",num1,num2)
	//交换后的两个数num1 = 10,num2 = 20  // 不变 因为是值传递

 }