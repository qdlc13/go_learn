package main
import(
	"fmt"
)
//var 数组名 [数组大小]数据类型
func main(){
	//实现的功能：给出五个学生的成绩，求出成绩的总和，平均数：
    //给出五个学生的成绩：--->数组存储：
    //定义一个数组
	var scores [5]int
	//将成绩存入数组
	scores[0] = 95
	scores[1] = 94
	scores[2] = 67
	scores[3] = 89
	scores[4] = 97
	//求和：
    //定义一个变量专门接收成绩的和
	sum := 0
	for i := 0;i < len(scores); i++ {//i: 0,1,2,3,4
		sum += scores[i]
	}
	//平均数：
	avg := sum / len(scores)
	//输出
	fmt.Printf("成绩的总和为：%v,成绩的平均数为：%v",sum,avg)
	//成绩的总和为：442,成绩的平均数为：88

}