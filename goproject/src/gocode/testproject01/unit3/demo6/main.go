package main
import(
	"fmt"
)
//switch后是一个表达式(即:常量值、变量、一个有返回值的函数等都可以)
//case后面的值如果是常量值(字面量)，则要求不能重复
//case后的各个值的数据类型，必须和 switch 的表达式数据类型一致
//case后面不需要带break 
//default语句不是必须的，位置也是随意的。

func main(){
//switch
	var score int = 67
	switch score/10 {
		case 10 :
			fmt.Println("您的等级为A级")
		case 9 :
			fmt.Println("您的等级为A级")
		case 8 :
			fmt.Println("您的等级为B级")
		case 7 , 6 , 5 :
			fmt.Println("您的等级为C级")
		default :
			fmt.Println("您的成绩有误")
	}
	//您的等级为C级
	//switch后也可以不带表达式，当做if分支来使用
	var a int32 = 2
	switch{
		case a == 1 :
			fmt.Println("aaa")
		default :
			fmt.Println("ddd")
		case a == 2 :
			fmt.Println("bbb")
		case a == 3 :
			fmt.Println("ccc")
		
	}
	//bbb
	//switch后也可以直接声明/定义一个变量，分号结束，不推荐
	switch b := 7 ; {
		case b > 6 :
			fmt.Println("大于6")
		case b <= 6 :
			fmt.Println("不大于6")
	}
	//大于6

	//switch穿透，利用fallthrough关键字，如果在case语句块后增加fallthrough ,则会继续执行下一个case,也叫switch穿透。
	switch b := 7 ; {
		case b > 6 :
			fmt.Println("大于6")
			fallthrough
		case b < 6 :
			fmt.Println("小于6")
		default :
			fmt.Println("等于6")
	}
	//大于6
	//小于6

}