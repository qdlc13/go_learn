package main
import(
	"fmt"
)

// func   函数名（形参列表)（返回值类型列表）{
// 	执行语句..
// 	return + 返回值列表
// }
/* 
	函数名：
	遵循标识符命名规范:见名知意 addNum,驼峰命名addNum  
	首字母不能是数字
	首字母大写该函数可以被本包文件和其它包文件使用(类似public)
	首学母小写只能被本包文件使用，其它包文件不能使用(类似private)
*/
//自定函数

func cal (num1 int,num2 int) (int) {   //返回值如果一个可以不加括号例如 int
	var sum int = 0
	sum += num1
	sum += num2
	return sum         //可以返回多个例如 return sum,res
}
func cal2 (num1 int,num2 int) (int,int){
	var sum int = num1 + num2
	var res int = num1 - num2
	return sum,res

}
func main(){
	sum := cal(10,20)
	fmt.Println(sum) //30

	var num3 int = 30
	var num4 int = 50

	sum1 := cal(num3,num4)
	fmt.Println(sum1) //80

	n1,n2 := cal2(10,20)  //n1,_ := cal2(10,20) 有返回值不想接收用_省略
	fmt.Println(n1,n2)  //30  -10

}