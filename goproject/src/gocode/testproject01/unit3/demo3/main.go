package main
import(
	"fmt"
)
func main(){
	var age int = 18
	fmt.Println("age对应存储空间的地址为：",&age)//age对应存储空间的地址为： 0xc000016088

	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr这个指针指向的具体数值为：",*ptr)//ptr这个指针指向的具体数值为： 18
	fmt.Println("ptr这个指针变量对应存储空间的地址为：",&ptr)//ptr这个指针变量对应存储空间的地址为： 0xc00000a030
}