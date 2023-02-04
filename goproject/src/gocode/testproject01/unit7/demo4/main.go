package main
import(
	"fmt"
)
//结构体类型是值类型，方法调用中，遵守值类型的传递机制，是值拷贝处理方式
//如果希望在方法中改变结构体变量的值可以通过结构体指针的方式处理
type Person struct{
	Name string
}
func (p *Person) test(){  //只要这里定义成为指针类型编译器就可以优化
	(*p).Name = "哥哥"
	fmt.Println((*p).Name)  //可以写p.Name底层自动转化
}
func main() {
	var p Person
	p.Name = "姐姐"
	fmt.Printf("p的地址是%p\n",&p)//p的地址是0xc00004e260
	(&p).test()	//哥哥  可以简化写p.test()底层自动转化
	fmt.Println(p.Name)//哥哥
	
}