package main
import(
	"fmt"
)
func main() {
	//定义切片方式一 定义一个切片，然后让切片去引用一个已经创建好的数组。
	var intarr [6]int = [6]int{4,3,87,34,77,3}
	slice := intarr[1:3]//等价 var slice []int = intarr[1:3]
	fmt.Println("切片一",slice)//切片一 [3 87]
	fmt.Println("切片长度",len(slice))//切片长度 2
	fmt.Println("切片容量",cap(slice))//切片容量 5


	//方式二
	//通过make内置函数来创建切片。基本语法: var切片名[type = make([], len,[cap])
	//参数 切片类型 切片长度 切片容量
	//PS : make底层创建一个数组，对外不可见，所以不可以直接操作这个数组，
	//要通过slice去间接的访问各个元素，不可以直接对数组进行维护/操作
	slice2 := make([]int,4,20)
	fmt.Println("切片长度",len(slice2))//切片长度 4
	fmt.Println("切片容量",cap(slice2))//切片容量 20
	slice2[1] = 6
	slice2[3] = 4
	fmt.Println(slice2)//[0 6 0 4]

	//方式3：定一个切片，直接就指定具体数组，使用原理类似make的方式
	slice3 := []int{2,0,54,3,6,5,4,34,34}
	fmt.Println(slice3)//[2 0 54 3 6 5 4 34 34]
	fmt.Println("切片长度",len(slice3))//切片长度 9
	fmt.Println("切片容量",cap(slice3))//切片容量 9

	//切片遍历
	//普通for循环
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("slice2[%v] = %v \n",i,slice2[i])
	}
	/* 
		slice2[0] = 0
		slice2[1] = 6
		slice2[2] = 0
		slice2[3] = 4	
	*/

	//for-range
	for i,v := range slice2 {
		fmt.Printf("slice2[%v] = %v \n",i,v)
	}

	/* 
		slice2[0] = 0
		slice2[1] = 6
		slice2[2] = 0
		slice2[3] = 4	
	*/

}