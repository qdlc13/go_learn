package main
import(
	"fmt"
)
//数组优点：访问/查询/读取 速度快
func main() {
	//声明数组
	var arr [3]int16
	//获取数组长度
	fmt.Println(len(arr))//3
	//打印数组:
	fmt.Println(arr)//[0 0 0]
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	//int16 两个字节吧    //数组每个空间占用的字节数取决于数组类型
	//证明arr中储存的地址值:
	fmt.Printf("arr的地址是:%p\n",&arr)//arr的地址是:0xc000016088
	//第一个空间的地址:
	fmt.Printf("arr的地址是:%p\n",&arr[0])//arr的地址是:0xc000016088
	//第二个空间的地址:
	fmt.Printf("arr的地址是:%p\n",&arr[1])//arr的地址是:0xc00001608a
	//第三个空间的地址:
	fmt.Printf("arr的地址是:%p\n",&arr[2])//arr的地址是:0xc00001608c




}