package main
import(
	"fmt"
)
//基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值。
func test(num int) {
	num = 50
	fmt.Println("test中num",num) //如此打印""中和后面的变量有空格
}

//以值传递方式的数据类型，如果希望在函数内的变量能修改函数外的变量，可以传入变量的地址&，
//函数内以指针的方式操作变量。从效果来看类似引用传递。
func test2(num *int){
	*num = 39
}
func main() {
	var num int = 10
	test(num)
	fmt.Println("main中num",num)
	/* 
		test中num 50
		main中num 10
	*/

	
	var num2 int = 2
	fmt.Println(&num2) //0xc000016088
	test2(&num2)
	fmt.Println(num2) //39

}