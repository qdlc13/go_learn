package main
import(
	"fmt"
)
/* 
【1】普通for循环
【2】键值循环
（键值循环） for range结构是Go语言特有的一种的迭代结构，在许多情况下都非常有用，for range 可以遍历数组、切片、字符串、map 及通道，for range 语法上类似于其它语言中的 foreach 语句，一般形式为：
for key, val := range coll {
    ...
}

注意：
（1）coll就是你要的数组
（2）每次遍历得到的索引用key接收，每次遍历得到的索引位置上的值用val
（3）key、value的名字随便起名  k、v   key、value  
（4）key、value属于在这个循环中的局部变量
（5）你想忽略某个值：用_就可以了：
*/
func main() {
	var scores [5]int
	//输入
	for i := 0; i < len(scores); i++ {
		fmt.Printf("输入第%d学生的成绩", i + 1)
		fmt.Scanln(&scores[i])
	}
	//展示
	//for循环
	for i := 0; i < len(scores); i++ {
		fmt.Printf("第%d个学生的成绩是：%d\n",i+1,scores[i])
	}
	fmt.Println("============================================================")
	//for-range循环
	for key,val := range scores {
		fmt.Printf("第%d个学生成绩是%d\n",key + 1,val)
	}
	for _,val := range scores {   //
		fmt.Printf("学生成绩是%d\n",val)
	}
	for key,_ := range scores {   //
		fmt.Printf("学生编号是%d\n",key+1)
	}
}
/* 
	输入第1学生的成绩23 
	输入第2学生的成绩54 
	输入第3学生的成绩98 
	输入第4学生的成绩15 
	输入第5学生的成绩88
	第1个学生的成绩是：23
	第2个学生的成绩是：54
	第3个学生的成绩是：98
	第4个学生的成绩是：15
	第5个学生的成绩是：88
	============================================================
	第1个学生成绩是23
	第2个学生成绩是54
	第3个学生成绩是98
	第4个学生成绩是15
	第5个学生成绩是88
	学生成绩是23
	学生成绩是54
	学生成绩是98
	学生成绩是15
	学生成绩是88
	学生编号是1
	学生编号是2
	学生编号是3
	学生编号是4
	学生编号是5
*/