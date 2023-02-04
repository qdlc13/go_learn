package main 
import(
	"fmt"
)
/* 
	golang支持多继承，如果一个结构体嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套结构体的字段
	和方法，从而实现多继承，但是为了保证代码的简介性，尽量不要使用多继承，很多语言去除了多继承
	go保留了
*/
type A struct{
	a int
	b string
}
type B struct{
	c int
	d string
	a int
}
//结构体的匿名字段可以是基本数据类型
type C struct{
	A
	B
	int
}
//嵌入匿名结构体指针也是可以的
type C1 struct{
	*A
	*B
	int
}
//结构体的字段可以是结构体类型的(组合模式)
type D struct{
	i int
	a A //这就不属于匿名(继承)了 不是继承关系

}

func main(){
	//C实例
	//等价c := C{A{12,"rrr"},B{46,"eee",50},8888}
	c := C{
			A{
				a : 12,
				b : "rrr",},
			B{	
				c : 46,
				d : "eee",
				a : 50,},
			8888,}


	fmt.Println(c)//{{12 rrr} {46 eee 50} 8888}

	fmt.Println(c.int)//8888
	fmt.Println(c.b)//rrr
	fmt.Println(c.d)//eee
	fmt.Println(c.A.a)//12
	fmt.Println(c.B.a)//50
	//如嵌入的匿名结构体有相同的字段和方法名，则访问时需要通过匿名结构体类型名区分
	/* fmt.Println(c.a) 相同名报错
	.\main.go:32:16: ambiguous selector c.a 
	*/

	//嵌入匿名结构体指针也是可以的
	c1 := C1{&A{10,"FFF"},&B{44,"UUUU",888},777}
	fmt.Println(c1)//{0xc000008078 0xc0000503e0 777}
	fmt.Println(*(c1.A))//{10 FFF} 等价fmt.Println(*c1.A))

	//组合模式
	d := D{29,A{23,"AAA"}}
	fmt.Println(d)//{29 {23 AAA}}
	fmt.Println(d.a.a)//23
	fmt.Println(d.a.b)//AAA
}