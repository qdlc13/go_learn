package main
import(
	"fmt"
	/* "unsafe" 引用没有使用的包
		# command-line-arguments
		.\main.go:4:2: imported and not used: "unsafe" 
	*/
)
func main(){
	//浮点类型 float32 4字节 float64 8字节
	var num1 float32 = 3.14
	var num2 float32 = -3.14
	fmt.Println(num1,num2) //3.14 -3.14

	//科学计数法 e E 均可
	var num3 float32 = 314e-2 
	var num4 float32 = 314E+2
	fmt.Println(num3,num4) //3.14 31400

	var num5 float32 = 314E+2
	var num6 float64 = 314E+2
	fmt.Println(num5,num6)    //31400 31400

	//浮点数可能会有精度的损失，所以通常情况下，建议你使用：float64 
	var num7 float32 = 256.000000916
	fmt.Println(num7) //256
	var num8 float64 = 256.000000916
	fmt.Println(num8) //256.000000916


	/* 
		bool //默认值为false
		string //默认值为空字符串
		int int8 int16 int32 int64 //默认值为0
		uint uint8 uint16 uint32 uint64 uintptr //默认值为0
		byte // uint8 的别名
		rune // int32 的别名
		float32 float64 //默认值为0
		complex64 complex128 //默认值为0
	*/

	//golang 默认浮点是float64 整数是int64
	var num9 = 3.17
	var num10 = 5
	fmt.Printf("num9,num10的默认类型是 %T %T",num9,num10)//num9,num10的默认类型是 float64 int


}