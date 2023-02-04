package main
import(
	"fmt"
)
func main() {
	//定义一个数组
	//长度属于类型的一部分
	var arr1 = [3]int{3,6,9}
	fmt.Printf("数组的类型是%T\n",arr1)
	//数组的类型是[3]int

	//Go中数组属值类型，在默认情况下是值传递，因此会进行值拷贝
	test(arr1)
	fmt.Println(arr1)//[3 6 9]

	//在其它函数中，去修改原来的数组，可以使用引用传递(指针方式)。
	test2(&arr1)
	fmt.Println(arr1) //[3 88 9]

}

func test(arr [3]int) {
	arr[1] = 88
}

func test2(arr *[3]int) {
	(*arr)[1] = 88
}