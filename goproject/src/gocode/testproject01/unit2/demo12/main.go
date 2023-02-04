package main
import(
	"fmt"
	"strconv"
)
func main(){
	//基本数据类型转字符串string
	var n1 int = 18
	////参数：第一个参数必须转为int64类型 ，第二个参数指定字面值的进制形式为十进制
	var s1 string = strconv.FormatInt(int64(n1),10)
	fmt.Printf("s1对应的类型是: %T , s1 = %q \n",s1,s1) //s1对应的类型是: string , s1 = "18"

	var n2 float64 = 5.21
	//第二个参数：'f'（-ddd.dddd）  第三个参数：9 保留小数点后面9位  第四个参数：表示这个小数是float64类型
	var s2 string = strconv.FormatFloat(n2,'f',9,64)
	fmt.Printf("s2对应的类型是: %T , s2 = %q \n",s2,s2) //s2对应的类型是: string , s2 = "5.210000000"

	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3对应的类型是: %T , s3 = %q \n",s3,s3) //s3对应的类型是: string , s3 = "true"

}