package dbutils
//同一文件夹（级别）不同文件的包名声明必须一致
import(
	"fmt"
)
func GetConn() { //首字母小写其他包不能访问
	fmt.Println("执行了dbutils包下的getConn()函数")
}