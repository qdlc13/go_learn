package dbutils
//同一文件夹（级别）不同文件的包名声明必须一致

//一个目录下不能有重复的函数
//同一个包下不能有同名函数
/* 
	# gocode/testproject01/unit4/demo9/crm/dbutils
	..\dbutils\utils.go:6:6: GetConn redeclared in this block
    ..\dbutils\dbutils.go:5:6: other declaration of GetConn
*/


// import(
// 	"fmt"
// )
// func GetConn() { //首字母小写其他包不能访问
// 	fmt.Println("执行了dbutils包下的getConn()函数")
// }

