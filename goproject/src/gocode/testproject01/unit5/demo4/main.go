package main
import(
	"fmt"
)
func main() {
	//数组定义赋值
	//第一种
	var arr1 [3]int = [3]int{3,6,9}
	fmt.Println(arr1)//[3 6 9]
	//第二种
	var arr2 = [3]int{1,2,44}
	fmt.Println(arr2)//[1 2 44]
	//第三种
	var arr3 = [...]int{2,3,5,7,43}
	fmt.Println(arr3)//[2 3 5 7 43]
	//第四种
	var arr4 = [...]int{2:77,0:322,1:88,7:23}//[322 88 77 0 0 0 0 23]
	fmt.Println(arr4)
}