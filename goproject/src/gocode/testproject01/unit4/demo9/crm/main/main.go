package main//package 进行包的声明 建议包的声明和所在文件夹同名
//main包是程序入口包 一般main函数会在这个包下

/* 	注意：
	import导入语句通常放在文件开头包声明语句的下面。
	导入的包名需要使用双引号包裹起来。
	包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔。
	需要配置一个环境变量：GOPATH 
*/
import(
	"fmt"
	//导入包文件夹的路径
	//可以给包取别名，取别名后，原来的包名就不能使用了 mypackage别名
	mypackage "gocode/testproject01/unit4/demo9/crm/dbutils"
	"gocode/testproject01/unit4/demo9/crm/calutils"

	
)
//我们不可能把所有的函数放在同一个源文件中，可以分门别类的把函数放在不同的原文件中 
//解决同名问题：两个人都想定义一个同名的函数，
//在同一个文件中是不可以定义相同名字的函数的。此时可以用包来区分 
func main() {
	fmt.Println("main函数执行")
	/* 
		省略dbutils.会引起错误
		导入但未使用
		imported and not used: "gocode/testproject01/unit4/demo9/crm/dbutils"
	*/

	//别名
	mypackage.GetConn() //函数调用要定位到所在的包 包名.方法名
	aaa.Add()
	/* 
		main函数执行
		执行了dbutils包下的getConn()函数
		add函数被执行了
	*/

}