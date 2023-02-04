package main
import(
	"fmt"
	"gocode/testproject01/unit7/demo10/model"

)
func main(){
	//跨包创建结构体Student实例
	//var s model.Student = model.Student{"哈哈",54} 等价
	s := model.Student{"哈哈",54}
	fmt.Println(s)//{哈哈 54}
	//工厂模式
	t := model.NewTeacher("老师",49)
	fmt.Println(*t)//{老师 49}

}