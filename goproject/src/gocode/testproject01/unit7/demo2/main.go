package main
import(
	"fmt"
)
type Student struct {
	Age int
	//两个结构体能类型转换 需要有完全一致的字段 age int也不可以
	
}
type Person struct {
	Age int
}
//结构体进行type重新定义(相当于取别名)，Golang认为是新的数据类型，但是相互间可以强转
type Stu Student


func main() {
	var s Student = Student{10}
    var p Person = Person{10}
	s = Student(p)
    fmt.Println(s)//{10}
    fmt.Println(p)//{10}

	var s1 Student = Student{19}
    var s2 Stu = Stu{19}
	s1 = Student(s2)
    fmt.Println(s1)//{19}
    fmt.Println(s2)//{19}


}