package main
import(
	"fmt"
)
//方法和函数区别
/* 
	函数：参数类型对应是什么就要传入什么
	方法：接收者是值类型可以传入指针类型，接收者是指针类型可以传入值类型
*/
type Student struct{
	Name string
}
//方法
func (s Student) test01(){
	fmt.Println(s.Name)
}
func (s *Student) test02(){
	fmt.Println((*s).Name)
}
//函数
func method01(s Student){
	fmt.Println(s.Name)
}
func method02(s *Student){
	fmt.Println((*s).Name)
}
func main() {
	var s Student = Student{"莉莉"}

	method01(s)//莉莉  //不能method01(&s)
	method02(&s)//莉莉 //不能method02(s)

	s.test01()//莉莉 
	(&s).test01()//莉莉 //虽然传入指针类型但是还是按照值传递方式 

	(&s).test02()//莉莉
	s.test02()//莉莉  等价 底层还是按照指针传递方式
	

}