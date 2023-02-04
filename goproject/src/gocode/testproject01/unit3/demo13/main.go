package main
import(
	"fmt"
)
func main(){
	for i := 1 ; i <= 100 ; i++ {
		fmt.Println(i)
		if i == 14 {
				return //结束当前的函数
		}
	}
	fmt.Println("hello golang")  //不执行
	/* 
		1
		2
		3
		4
		5
		6
		7
		8
		9
		10
		11
		12
		13
		14

	*/

}