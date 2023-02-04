package main
import(
	"fmt"
)
//闭包就是一个函数和与其相关的引用环境组合的一个整体
//函数功能：求和
//函数的名字：getSum 参数为空
//getSum函数返回值为一个函数，这个函数的参数是一个int类型的参数，返回值也是int类型

func getSum() func (int) int {
	var sum int = 0
	return func (num int) int {
		sum += num
		return sum
	}
}
//闭包：返回的匿名函数+匿名函数以外的变量sum
//匿名函数中引用的那个变量会一直保存在内存中，可以一直使用 sum
/* 
	闭包的本质：
	闭包本质依旧是一个匿名函数，只是这个函数引入外界的变量/参数
	匿名函数+引用的变量/参数 = 闭包

	特点：
	（1）返回的是一个匿名函数，但是这个匿名函数引用到函数外的变量/参数 
	,因此这个匿名函数就和变量/参数形成一个整体，构成闭包。
	（2）闭包中使用的变量/参数会一直保存在内存中，所以会一直使用---》
	意味着闭包不可滥用（对内存消耗大）
*/

func main() {
	//此时f不看做getSum闭包函数 而是看作getSum返回的匿名函数
	f := getSum() 

	//闭包：返回的匿名函数+匿名函数以外的变量num
	fmt.Println(f(1))//1
	fmt.Println(f(2))//3
	fmt.Println(f(3))//6
	fmt.Println(f(4))//10


	//全局变量替换闭包
	fmt.Println(testSum(1))//1
	fmt.Println(testSum(2))//3
	fmt.Println(testSum(3))//6
	fmt.Println(testSum(4))//10
}

//不使用闭包的时候：我想保留的值，不可以反复使用
//闭包应用场景：闭包可以保留上次引用的某个值，我们传入一次就可以反复使用了

var t int = 0;
func testSum(num int) int {
	t += num
	return t

}