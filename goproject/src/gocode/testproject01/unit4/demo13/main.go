package main
import(
	"fmt"
)
//defer关键字的作用：
//在函数中，程序员经常需要创建资源，为了在函数执行完毕后，及时的释放资源，
//Go的设计者提供defer关键字
func main() {
	fmt.Println(add(20,90))
	fmt.Println("=================================================")
	fmt.Println(add2(40,60))
}
func add(num1 int, num2 int) int {
	//go中遇到defer关键字 不会立刻执行defer后的语句
	//而是将defer后的语句压入一个栈中，然后继续执行函数后面的语句
	defer fmt.Println("num1=",num1)
	defer fmt.Println("num2=",num2)

	//栈先进后出
	//在函数执行完毕后，从栈中取出语句开始执行,按照先进后出的顺序执行
	var sum int = num1 + num2
	fmt.Println("sum=",sum)
	return sum
}

/* 
	sum= 110
	num2= 90
	num1= 20
	110
*/

func add2(num1 int, num2 int) int {
	//遇到defer关键字，会将后面的代码语句压入栈中，
	//也会将相关的值同时拷贝入栈中，不会随着函数后面的变化而变化
	defer fmt.Println("num1=",num1)//40
	defer fmt.Println("num2=",num2)//60
	
	num1 += 90 //130
	num2 += 50 //110
	var sum int = num1 + num2
	fmt.Println("sum=",sum) //240
	return sum
}
/* 
	sum= 240
	num2= 60
	num1= 40
	240
*/

/* 	defer应用场景
	defer+资源回收代码
	比如你想关闭某个使用的资源，在使用的时候直接随手defer，
	因为defer有延迟执行机制（函数执行完毕再执行defer压入栈的语句），
	所以你用完随手写了关闭，比较省心，省事
 */