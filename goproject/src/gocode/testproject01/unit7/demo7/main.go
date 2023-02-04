package main
import(
	"fmt"
)
//方法和函数区别
/* 
	方法需要绑定指定数据类型
	函数不需要绑定数据类型
	调用方式不用：
	函数：函数名(实参列表)
	方法：变量.方法名(实参列表)
*/
type Student struct{
	Name string
	Age int
}
//方法
func (s Student) test01(){
	fmt.Println(s.Name)
}
//函数
func method01(s Student){
	fmt.Println(s.Name)
}
func main() {
	var s Student = Student{"莉莉",20}//等价s := Student{"莉莉",20}
	
	method01(s)//莉莉
	s.test01()//莉莉

}