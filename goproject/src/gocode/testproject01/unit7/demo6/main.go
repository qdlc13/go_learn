package main
import(
	"fmt"
)
type Student struct{
	Name string
	Age int
}
func (s *Student) String() string {
	str := fmt.Sprintf("Name = %v , Age = %v",(*s).Name,(*s).Age)
	return str
}
//如果一个类型实现了String方法，那么fmt.Println默认调用这个变量的String()进行输出
//以后定义结构体的话，常定义String()作为输出结构体信息的方法，在fmt.Println会自动调用
func main() {
	stu := Student{
		Name : "煮酒",
		Age : 20 ,
	}
	fmt.Println(stu)//{煮酒 20}
	//传入地址,如果绑定了String方法会自动调用
	fmt.Println(&stu)//Name = 煮酒 , Age = 20

}