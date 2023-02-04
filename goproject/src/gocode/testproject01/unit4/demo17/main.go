package main
import(
	"fmt"
)
func main() {
	//程序中出现错误/恐慌以后，程序被中断，无法继续执行。
	test()
	fmt.Println("上面除法正常执行")//未执行
	fmt.Println("执行新的逻辑")//未执行
}
func test() {
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}
/* panic(恐慌): runtime error: integer divide by zero

goroutine 1 [running]:
main.test()
        C:/Users/qdliu/Desktop/goproject/src/gocode/testproject01/unit4/demo17/main.go:13 +0x11
main.main()
        C:/Users/qdliu/Desktop/goproject/src/gocode/testproject01/unit4/demo17/main.go:6 +0x1d
exit status 2 
*/
