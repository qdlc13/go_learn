package main
import(
	"fmt"
	"gocode/testproject01/unit4/demo10/testutils"
)
var num int = test()
func test() int {
	fmt.Println("main.go执行test函数")
	return 10
}
//init函数：初始化函数，可以用来进行一些初始化的操作
//每一个源文件都可以包含一个init函数，该函数会在main函数执行前，
//被Go运行框架调用。
func init(){
	fmt.Println("main.go执行init函数")
}
func main(){
	fmt.Println("main.go执行main函数")
	fmt.Println("Age=",testutils.Age,"Sex=",testutils.Sex,"Name=",testutils.Name)
}
/*  执行过程
	testutils中init函数执行
	main.go执行test函数
	main.go执行init函数
	main.go执行main函数
	Age= 20 Sex= 女 Name= 哈哈
*/

//执行过程 先执行引入包的变量定义->引入包的init函数->main包的变量定义->main包的init函数->main函数