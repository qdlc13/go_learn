package main
import(
	"fmt"
	"strconv"
)
func main(){
	//string转基本数据类型
	var s1 string = "true"
	var b bool
	/*报错strconv.ParseBool返回值有两个  func ParseBool(str string) (value bool, err error)
	 b = strconv.ParseBool(s1)
	 # command-line-arguments
	.\main.go:10:6: assignment mismatch: 
	1 variable but strconv.ParseBool returns 2 values
	 */
	//ParseBool这个函数的返回值有两个：(value bool, err error)
    //value就是我们得到的布尔类型的数据，err出现的错误
    //我们只关注得到的布尔类型的数据，err可以用_直接忽略
	b , _ = strconv.ParseBool(s1)
	fmt.Printf("b的类型是：%T,b=%v \n",b,b) //b的类型是：bool,b=true

	//string ->int64
	var s2 string = "19"
	var num1 int64
	num1 , _ =strconv.ParseInt(s2,10,64)
	fmt.Printf("num1类型是：%T,num1=%v \n",num1,num1) //num1类型是：int64,num1=19 

	//string ->float32/64
	var s3 string = "3.221"
	var f1 float64
	f1 , _ =strconv.ParseFloat(s3,64)
	fmt.Printf("f1类型是：%T,f1=%v \n",f1,f1)//f1类型是：float64,f1=3.221

	//注意：string向基本数据类型转换的时候，一定要确保string类型能够转成有效的数据类型，否则最后得到的结果就是按照对应类型的默认值输出
	var s4 string = "golang"
	var b1 bool
	b1 , _ = strconv.ParseBool(s4)
	fmt.Printf("b1的类型是：%T,b1=%v \n",b1,b1) //b1的类型是：bool,b1=false

	var s5 string = "golang"
    var num2 int64
    num2,_ = strconv.ParseInt(s5,10,64)
    fmt.Printf("num2的类型是：%T,num2=%v \n",num2,num2)//num2的类型是：int64,num2=0

}