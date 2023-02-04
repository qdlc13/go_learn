package testutils
import(
	"fmt"
)
var Age int
var Sex string
var Name string
//init函数初始化变量
func init() {
	fmt.Println("testutils中init函数执行")
	Age = 20
	Sex = "女"
	Name = "哈哈"

}