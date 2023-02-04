package main
import "fmt"
//全局变量定义
//n7,n8 := "quanju",19 "似乎不可以"
var n7,n8 = "quanju",19
//设计者认为上述复杂 使用一次性声明
var(
	n9 = 200
	n10 = "lc"
)

func main(){
	//{}中的为局部变量
	//变量使用方式
	var num int = 18
	fmt.Println(num)   //18
	//指定类型不赋值使用默认值
	var num2 int
	fmt.Println(num2)  //0
	//不写类型自动根据值推断
	var num3 = "tom"
	fmt.Println(num3)  //tom
	//省略var关键字 定义并赋值推断 不能写为 =
	sex := "男"  //只适用于局部变量
	fmt.Println(sex) //"男"

	fmt.Println("---------------------------------")

	//多变量声明
	var n1,n2,n3 int
	fmt.Println(n1,n2,n3) //0 0 0
	//自动推断类型
	var n4,name,n5 = 10,"jack",7.6 
	fmt.Println(n4,name,n5) //10 jack 7.6
	//定义并赋值
	n6,height := "nihao",200
	fmt.Println(n6,height) //nihao 200
	
	//全局
	fmt.Println(n7,n8) //quanju 19
	fmt.Println(n9,n10) //200 lc
}