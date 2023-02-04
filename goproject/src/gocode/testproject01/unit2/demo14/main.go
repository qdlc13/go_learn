package main
import(
	"fmt"
)
func main(){
	//指针
	var age int
	fmt.Println(&age) //0xc000016088

	//指针变量
	//定义一个指针变量：
    //var代表要声明一个变量
    //ptr 指针变量的名字
    //ptr对应的类型是：*int 是一个指针类型 （可以理解为 指向int类型的指针）
    //&age就是一个地址，是ptr变量的具体的值
	var ptr *int = &age
	fmt.Println(ptr) //0xc000016088
	fmt.Println("ptr本身这个存储空间地址是:",&ptr)//ptr本身这个存储空间地址是: 0xc00000a030

	/* 
		1.& 取内存地址
		2.* 根据地址取值
	*/
}