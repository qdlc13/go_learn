package main
import(
	"fmt"
)
func main(){
	//错误处理/捕获机制：
	//go中追求代码优雅，引入机制：defer+recover机制处理错误
	test()
	fmt.Println("上面除法正常执行")//未执行
	fmt.Println("执行新的逻辑")//未执行
}
func test() {

	//利用defer + recover捕获错误，defer后加上匿名函数调用
	//recover函数见文档
	defer func() {
		//调用recover内置函数捕获错误
		err := recover()
		//如果没有捕获错误返回值是nil
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println("err是",err)
		}
	}()//加()调用这个匿名函数

	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}
//提高系统健壮性
/* 
错误已经捕获
err是 runtime error: integer divide by zero
上面除法正常执行
执行新的逻辑
*/