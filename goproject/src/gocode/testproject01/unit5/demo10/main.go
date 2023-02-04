package main
import(
	"fmt"
)
func main() {
	var intarr [6]int = [6]int{1,4,5,7,3,6}
	//切片
	slice := intarr[1:4]
	fmt.Println(slice)//[4 5 7]
	fmt.Println(len(slice)) //3

	//追加元素
	slice2 := append(slice,75,33,3)
	fmt.Println(slice2)//[4 5 7 75 33 3]
	/* 
		底层原理：
        1.底层追加元素的时候对数组进行扩容，老数组扩容为新数组：
        2.创建一个新数组，将老数组中的4,5,7复制到新数组中，在新数组中追加75,33,3
        3.slice2 底层数组的指向 指向的是新数组 
        4.往往我们在使用追加的时候其实想要做的效果给slice追加：
	*/

	slice = append(slice,66,99)
	fmt.Println(slice)//[4 5 7 66 99] 
	fmt.Println(intarr)//[1 4 5 7 66 99] 数组后面被覆盖了

	slice[0] = 1000
	fmt.Println(slice)//[1000 5 7 66 99] 
	fmt.Println(intarr)//[1 1000 5 7 66 99]
	fmt.Println(slice2)//[4 5 7 75 33 3] slice2没有改变因为slice2底层维护的是新数组
	//5.底层的新数组 不能直接维护，需要通过切片间接维护操作。

	//可以通过append函数将切片追加给切片
	slice3 := []int{666,777}
	slice = append(slice,slice3...)
	fmt.Println(slice)//[1000 5 7 66 99 666 777]
	slice[0]=0
	fmt.Println(slice)//[0 5 7 66 99 666 777]
	fmt.Println(intarr)//[1 1000 5 7 66 99] 没有改变因为slice也维护新数组了

	//切片的拷贝
	var a []int = []int{93,84,37,66,18}
	var b []int = make([]int,10)//容量省略
	//拷贝吧a给b 拷贝操作修改任意值互不影响
	copy(b,a)
	fmt.Println(b)//[93 84 37 66 18 0 0 0 0 0]


}