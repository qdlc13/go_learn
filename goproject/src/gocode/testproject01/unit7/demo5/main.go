package main
import(
	"fmt"
)
//golang 中方法作用在指定数据类型上，和指定数据类型绑定
//所以自定义类型都可以有方法不仅仅是struct，int，float32
//等也可以有方法
type Intege int //简单数据类型必须先起别名在添加方法
//方法首字母大写其他包和本包都可以访问到，首字母小写只能在本包访问
func (i Intege) print(){
	i = 289
	fmt.Println("i = ",i)
}
func (i *Intege) change(){
	*i = 9
	fmt.Println("i = ",*i)
}
func main() {
	var i Intege = 29
	i.print() //i =  289
	fmt.Println(i)//29

	(&i).change()//i =  9
	fmt.Println(i)//9
}