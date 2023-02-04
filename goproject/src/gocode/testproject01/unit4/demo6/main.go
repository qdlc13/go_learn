package main
import(
	"fmt"
)
//在Go中，函数也是一种数据类型，
//可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用。
func test(num int) {
	fmt.Println(num)
}

//函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
//（把函数本身当做一种数据类型）
func test2(num int ,num1 float32, testFun func(int)) {
	testFun(3)
	fmt.Println(num,num1)
}

func main(){
	//函数也是一种数据类型，可以赋值给一个变量	
	a := test
	fmt.Printf("a的类型是%T,test的类型是%T\n",a,test)
	//a的类型是func(int),test的类型是func(int)

	//通过该变量可以对函数调用
	a(10) //等价于 test(10) //10

	//调用test02函数：
	test2(10,3.19,test) 
	/* 	3
		10 3.19 
	*/
	test2(10,3.19,a)   
	/* 	3
		10 3.19 
	*/


}