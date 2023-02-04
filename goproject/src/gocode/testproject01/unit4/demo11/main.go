package main
import(
	"fmt"
)
var Func01 = func (num1 int, num2 int) int {
	return num1 * num2
}
//Go支持匿名函数，如果我们某个函数只是希望使用一次，可以考虑使用匿名函数
func main() {
	//在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次（用的多）
	//定义匿名函数 定义的同时调用
	result := func (num1 int, num2 int) int {
		return num1 + num2
	}(10,20)
	fmt.Println(result)  //30

	//将匿名函数赋给一个变量(该变量就是函数变量了)，再通过该变量来调用匿名函数（用的少）
	//将匿名函数赋给一个变量，这个变量实际上就是函数类型的变量
	//sub 等价于匿名函数
	sub := func (num1 int, num2 int) int {
		return num1 - num2
	}
	//直接调用sub就是调用这个匿名函数了

	result1 := sub(30,40)
	fmt.Println(result1)//-10
	result2 := sub(30,40)
	fmt.Println(result2)//-10

	//如何让一个匿名函数，可以在整个程序中有效呢?
	//将匿名函数给一个全局变量就可以了

	result3 := Func01(3,4)
    fmt.Println(result3) //12

}