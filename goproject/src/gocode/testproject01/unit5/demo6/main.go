package main
import(
	"fmt"
)
func main() {
	//二维数组
	var arr [2][3]int16
	//默认值
	fmt.Println(arr)//[[0 0 0] [0 0 0]]

	//内存
	fmt.Printf("arr地址:%p\n",&arr)//arr地址:0xc0000160a0
	fmt.Printf("arr[0]地址:%p\n",&arr[0])//arr[0]地址:0xc0000160a0
	fmt.Printf("arr[0][0]地址:%p\n",&arr[0][0])//arr[0][0]地址:0xc0000160a0

	fmt.Printf("arr[0][1]地址:%p\n",&arr[0][1])//arr[0][1]地址:0xc0000160a2
	fmt.Printf("arr[0][2]地址:%p\n",&arr[0][2])//arr[0][2]地址:0xc0000160a4


	fmt.Printf("arr[1]地址:%p\n",&arr[1])//arr[1]地址:0xc0000160a6
	fmt.Printf("arr[1][0]地址:%p\n",&arr[1][0])//arr[1][0]地址:0xc0000160a6
	fmt.Printf("arr[1][1]地址:%p\n",&arr[1][1])//arr[1][0]地址:0xc0000160a8

	//赋值
	arr[0][0] = 58
	arr[1][1] = 99
	fmt.Println(arr)//[[58 0 0] [0 99 0]]

	//初始化 省略写法		var arr2 = [2][3]int{{2,3,4},{2,6,8}}
	var arr2 [2][3]int= [2][3]int{{2,3,4},{2,6,8}}
	fmt.Println(arr2)//[[2 3 4] [2 6 8]]

	//遍历二维数组
	//普通for循环
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Printf("arr2[%v][%v] = %v ",i,j,arr2[i][j])
		}
		fmt.Println()
	}
	/* 	
		arr2[0][0] = 2 arr2[0][1] = 3 arr2[0][2] = 4
		arr2[1][0] = 2 arr2[1][1] = 6 arr2[1][2] = 8
	*/
	//for-range
	for key,value := range arr2 {
		for k,v := range value {
			fmt.Printf("arr2[%v][%v] = %v ",key,k,v)
		}
		fmt.Println()
	}

	/* 
		arr2[0][0] = 2 arr2[0][1] = 3 arr2[0][2] = 4
		arr2[1][0] = 2 arr2[1][1] = 6 arr2[1][2] = 8
	*/

}