package main //main是程序的入口包 package编译和执行都不可以
//包名最好和目录保持一致
//package fmt 不可以和标准库不要冲突

/* 	注意：
	import导入语句通常放在文件开头包声明语句的下面。
	导入的包名需要使用双引号包裹起来。
	包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔。
	需要配置一个环境变量：GOPATH 
*/
import(
	"fmt"
	//绝对路径
	"gocode/testproject01/unit2/demo16/test"
)

func main(){
	
	var num int = 27 
	//驼峰命名法
	var stuNameDetail string = "lili"
	/* 
		如果变量名、函数名、常量名首字母大写，则可以被其他的包访问;
         如果首字母小写，则只能在本包中使用  （利用首字母大写小写完成权限控制）
	*/
	fmt.Println(stuNameDetail,num) //lili 27

	//输出学号
	//go env -w GO111MODULE=off 关闭了go mod 模式 后续需要使用
	//https://blog.csdn.net/qq_42647903/article/details/123253707
	fmt.Println(test.StuNo)
}