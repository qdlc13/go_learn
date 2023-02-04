package main
// import "fmt"
// import "unsafe" 等价
import(
	"fmt"
	"unsafe"
)
/*  
	rune类型是当需要处理中文、日文或其他复合字符时，需要用到的 等价int32
	byte是uint8的别名
 */
func main(){
	/* var num int8 = 1000    赋值越界
	# command-line-arguments
	.\main.go:5:17: cannot use 1000 (untyped int constant) as 
	int8 value in variable declaration (overflows)

	var num uint8 = -20     越界
	# command-line-arguments
	.\main.go:9:18: cannot use -20 (untyped int constant) as 
	uint8 value in variable declaration (overflows)
	*/
	
	var num uint16 = 8 //fmt.Printf把数据类型填充到 %T
	var num1 int8 = 3
	var num2 int16 = 3
	var num3 = 3 // 等价于 var num3 int64 = 3
	fmt.Printf("num的类型是：%T",num)  //num的类型是：uint16

	fmt.Println()

	fmt.Println(unsafe.Sizeof(num))  //字节数  2

	fmt.Println(unsafe.Sizeof(num1))  //字节数 1

	fmt.Println(unsafe.Sizeof(num2))  //字节数 2

	fmt.Println(unsafe.Sizeof(num3))  //字节数 8 
}