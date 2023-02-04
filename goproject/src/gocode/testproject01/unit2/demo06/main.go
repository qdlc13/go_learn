package main
import(
	"fmt"
	"unsafe"
)
func main(){
	//布尔值
	var flag1 bool = true
	fmt.Println(flag1) //true
	var flag2 bool = false
	fmt.Println(flag2) //false

	var flag3 bool = 5 < 9
	fmt.Println(flag3) //true

	fmt.Println(unsafe.Sizeof(flag3)) //字节数 1
	fmt.Printf("%d",unsafe.Sizeof(flag3)) //1
}