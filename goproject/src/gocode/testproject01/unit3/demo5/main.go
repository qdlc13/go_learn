package main
import(
	"fmt"
)
func main(){
	var count int = 100
	//单分支
	if count < 200 {
		fmt.Println("数量不足")//数量不足
	}

	//if后面一定要有空格，和条件表达式分隔开来
    //{}一定不能省略
    //条件表达式左右的()是建议省略的
    //在golang里，if后面可以并列的加入变量的定义：
	//sum在这个if之外不能使用
	if sum := 5;sum > 2 {
		fmt.Println("数量超过了") //数量超过
	}

	if sum1 := 3;sum1 > 3 {
		fmt.Println("分支a")
	}else{                        //else不能另起一行
		fmt.Println("分支b")      //分支b
	}
	//多分支
	var score int = 79
	if score >= 90 {
		fmt.Println("成绩为A级别")
	} else if score >= 80 {
		fmt.Println("成绩为B级别")
	} else {
		fmt.Println("成绩为C级别")//成绩为C级别
	}
	

	/* 报错 sum未定义 .\main.go:36:5: undefined: sum
	if sum > 10 {
		fmt.Println("aaa")
	} else if sum > 6{
		fmt.Println("bbb")
	} 
	*/
}
