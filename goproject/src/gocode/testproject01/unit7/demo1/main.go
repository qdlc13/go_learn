package main
import(
	"fmt"
)

/* 
	Golang语言面向对象编程说明：
（1）Golang也支持面向对象编程(OOP)，但是和传统的面向对象编程有区别，
并不是纯粹的面向对象语言。所以我们说Golang支持面向对象编程特性是比较准确的。
（2）Golang没有类(class)，Go语言的结构体(struct)和其它编程语言的类(class)有同等的地位，
你可以理解Golang是基于struct来实现OOP特性的。
（3）Golang面向对象编程非常简洁，去掉了传统OOP语言的方法重载、构造函数和析构函数、
隐藏的this指针等等
（4）Golang仍然有面向对象编程的继承，封装和多态的特性，
只是实现的方式和其它OOP语言不一样，比如继承:Golang没有extends关键字，继承是通过匿名字段来实现。 
*/

//结构体 (是值类型)
//定义老师结构体，将老师中的各个属性  统一放入结构体中管理：
type Teacher struct{
	//变量名字大写外界可以访问这个属性
	Name string
	Age int
	School string

}
func main() {
	var t1 Teacher
	fmt.Println(t1)//没有赋值 { 0 }
	t1.Name = "哈哈哈"
	t1.Age = 34
	t1.School = "清华大学"
	fmt.Println(t1)//{哈哈哈 34 清华大学}
	fmt.Println(t1.Age + 10) //44
	
	//其他创建方式
	var t2 Teacher = Teacher{"jack",34,"武汉大学"}
	fmt.Println(t2)//{jack 34 武汉大学}

	//创建教师结构体的实例对象变量
	var t3 *Teacher = new(Teacher)
	//t3是指针 t3指向的是地址 给这个地址指向的对象字段赋值
	(*t3).Name = "mike"
	(*t3).Age = 24
	//为符合程序员编程习惯go提供了简化的赋值方式
	t3.School = "扬州大学" //go编译器底层优化t.School转化为(*t).School赋值
	fmt.Println(*t3)//{mike 24 扬州大学}

	//方式四
	var t4 *Teacher = &Teacher{"green",28,"哈佛大学"}
	fmt.Println(*t4)//{green 28 哈佛大学}



}