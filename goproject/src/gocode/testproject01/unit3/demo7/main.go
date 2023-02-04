package main
import(
	"fmt"
)
func main(){

	var sum int = 0
	for i := 1 ; i <= 5 ; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for循环的语法格式：
        // for 初始表达式; 布尔表达式（条件判断）; 迭代因子 {
        // 	循环体;--》反复重复执行的内容
        // }
        // 注意：for的初始表达式 不能用var定义变量的形式，要用:=
        // 注意：for循环实际就是让程序员写代码的效率高了，但是底层该怎么执行还是怎么执行的，
		//底层效率没有提高，只是程序员写代码简洁了而已


		t := 1 //变量的初始化
        for t<=5 {//条件表达式。判断条件
                fmt.Println("你好 Golang")//循环体
                t++//迭代
        }

		var a int = 1 //变量的初始化
        for a = 2 ; a<=5 ; {//条件表达式。判断条件
                fmt.Println("你好")//循环体
				a++
        }

	// for ;; {
	// 		fmt.Println("你好 Golang")
	// }
}