package main
import(
	"fmt"
)
func main() {
	//切片注意点
	//切片定义后不可以直接使用，需要让其引用到一个数组，或者make一个空间供切片来使用
	var slice []int
	fmt.Println(slice) // []

	//不可越界
	slice2 := []int {0,1,2,3}  //[]空着是切片有...或者数字是数组
	fmt.Println(slice2) //[0 1 2 3]
	/* 
	   fmt.Println(slice2[4]) 
	   panic: runtime error: index out of range [4] with length 4
	*/ 

	var intarr [6]int = [6]int{1,4,7,2,5,8}
	//切片
	var slice3 []int = intarr[1:4]
	//切片可以继续切片
	slice4 := slice3[0:2]
	fmt.Println(slice3)//[4 7 2]
	fmt.Println(slice4)//[4 7]

	slice4[0] = 100
	//影响内存空间都改值
	fmt.Println(intarr)//[1 100 7 2 5 8]
	fmt.Println(slice3)//[100 7 2]
	fmt.Println(slice4)//[100 7]
}