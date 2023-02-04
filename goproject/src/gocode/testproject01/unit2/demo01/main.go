package main
import "fmt"
func main(){
	//变量声明
	var age int
	//赋值
	age = 18
	fmt.Println("age = ",age)

	//变量声明赋值
	var age2 int = 19
	fmt.Println("age2 = ",age2)

	/* var age int  重复定义出错
	# 	command-line-arguments
		.\main.go:14:6: age redeclared in this block
        .\main.go:5:6: other declaration of age
	*/

	/* var num int = 12.2  赋值类型不匹配出错
	# command-line-arguments   
     .\main.go:20:16: cannot use 12.2 (untyped float constant) 
	 as int value in variable declaration (truncated)
	*/
}