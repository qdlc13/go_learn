package main
import(
	"fmt"
)
func numChange(num1 int , num2 int) (int, int) {
	var sum int = num1 + num2 
	var sub int = num1 * num2
	return sum,sub
}

func test(num1 int, num2 int, testFunc func(int, int) (int, int)) {
	var _,t int = testFunc(num1,num2)
	fmt.Printf("测试成功乘积为%v\n",t)
}

//类型太复杂也可以用别名代替func(int, int) (int, int)
type myType func(int, int) (int, int)
func test2(num1 int, num2 int, testFunc myType) {
	var _,t int = testFunc(num1,num2)
	fmt.Printf("测试成功乘积为%v\n",t)
}
//go支持支持对函数返回值命名
//传统写法要求：返回值和返回值的类型对应，顺序不能差
func numTest1(num1 int , num2 int) (int, int) {
	result1 := num1 + num2 
	result2 := num1 - num2 
	return result1,result2
}
//升级写法：对函数返回值命名，里面顺序就无所谓了，顺序不用对应
func numTest2(num1 int , num2 int) (result1 int, result2 int) {
	result2 = num1 - num2 //注意顺序注意这里不要用 := 因为这里(result1 int, result2 int)已经声明
	result1 = num1 + num2 
	return 
}
func main() {

	a := numChange
	fmt.Printf("函数的类型%T,a的类型%T\n",numChange,a) //func(int, int) (int, int)
	
	test(2,3,a) //等价于test(2,3,numChange)//测试成功乘积为6

	//类型太复杂也可以用别名代替func(int, int) (int, int)
	test2(3,3,a) //等价于test(2,3,numChange)//测试成功乘积为9

	n1,n2 := numTest1(7,2)
	fmt.Printf("n1为和%v，n2为差%v\n",n1,n2) //n1为和9，n2为差5

	n1,n2 = numTest2(6,8)
	fmt.Printf("n1为和%v，n2为差%v",n1,n2) //n1为和14，n2为差-2

}