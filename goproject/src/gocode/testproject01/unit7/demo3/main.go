package main
import(
	"fmt"
)
//方法是作用在指定的数据类型上、和指定的数据类型绑定，因此自定义类型都可以有方法不仅仅是struct
//方法！=函数
//方法声明和调用格式
type A struct {
	Num int
}
//方法
//只能用A类型结构体调用 test方法与结构体A类型有绑定关系
//其他类型调用test方法报错
//结构体对象传入test方法中，值传递，和函数参数传递一致
func (a A) test() { 
	a.Num = 20
	fmt.Println(a.Num)
}
func main() {
	//调用
	var a A = A{13}
	a.test()//20
	fmt.Println(a.Num)//13 因为是值传递

}