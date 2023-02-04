package main
import(
	"fmt"
	/* 
	  下划线是空标识符
	  _"fmt"等价于把 //fmt 注释掉系统会忽略
	*/
)

/* 变量由字母数字下划线_组合构成 但是不能单独使用下划线 */
func main(){
	//指针细节
	//指针改变变量的值
	var num int = 10
	fmt.Println(num) //10

	/* var ptr *int = num   指针变量接受的一定是地址值
		# command-line-arguments
		.\main.go:11:17: cannot use num (variable of type int) 
		as type *int in variable declaration
	*/

	/* var ptr *float32 = &num 指针变量类型和变量类型必须匹配
		# command-line-arguments
		.\main.go:20:21: cannot use &num (value of type *int) 
		as type *float32 in variable declaration
	*/
	var ptr *int = &num 
	*ptr = 50
	fmt.Println(*ptr,num) //50 50
}