package main 
import(
	"fmt"
)
type Student struct{
	Name string
	Age int
}
func main() {
	//方式一 按照顺序赋值操作 缺点必须按照顺序
	var s1 Student = Student{"夏利",30}
	fmt.Println(s1)//{夏利 30}

	//方式二 按照指定类型赋值
	var s2 Student = Student{
		Age : 33,
		Name : "速度",
	}
	fmt.Println(s2)//{速度 33}

	//方式三 返回结构体指针类型
	var s3 *Student = &Student{"明明",32}
	fmt.Println(s3)//&{明明 32}
	fmt.Println(*s3)//{明明 32}

	var s4 *Student = &Student{
		Age : 13,
		Name : "哈哈",
	}
	fmt.Println(s4)//&{哈哈 13
	fmt.Println(*s4)//{哈哈 13}

}