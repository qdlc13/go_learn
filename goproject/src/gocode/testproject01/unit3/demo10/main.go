package main
import(
	"fmt"
)
func main(){
	//标签
	label:
	for i := 1 ; i <= 5 ; i++ {
		for j := 2 ; j <= 4 ; j++ {
			fmt.Printf("i: %v,j: %v\n",i,j)
			if i == 2 && j == 2 {
				break label
			}
		}
	}

	fmt.Println("-----ok") //执行
	/* 
		i: 1,j: 2
		i: 1,j: 3
		i: 1,j: 4
		i: 2,j: 2
		-----ok
	*/


}