package main
import(
	"fmt"
)
func main(){
	//continue
	for i := 1 ; i <= 100 ; i++ {
		if i % 6 == 0 {
			fmt.Println(i)
		}
	}

	for i := 1 ; i <= 100 ; i++ {
		if i % 6 != 0 {
			continue
		}
		fmt.Println(i)
	}
	//continue的作用是结束离它近的那个循环，继续离它近的那个循环

	for i := 1 ; i <= 6 ; i++ {
		for j := 2 ; j <= 4 ; j++ {
			if i == 2 && j == 2 {
				continue
			}
			fmt.Printf("i: %v, j: %v \n",i,j)
		}
	}

	/* 
		6
		12
		18
		24
		30
		36
		42
		48
		54
		60
		66
		72
		78
		84
		90
		96
		6
		12
		18
		24
		30
		36
		42
		48
		54
		60
		66
		72
		78
		84
		90
		96
		i: 1, j: 2
		i: 1, j: 3 
		i: 1, j: 4
		i: 2, j: 3
		i: 2, j: 4
		i: 3, j: 2
		i: 3, j: 3
		i: 3, j: 4
		i: 4, j: 2
		i: 4, j: 3
		i: 4, j: 4
		i: 5, j: 2
		i: 5, j: 3
		i: 5, j: 4
		i: 6, j: 2
		i: 6, j: 3
		i: 6, j: 4
	*/

	label:
	for i := 1 ; i <= 5 ; i++ {
		for j := 2 ; j <= 4 ; j++ {
			if i == 2 && j == 2 {
				continue label
			}
			fmt.Printf("i:%v, j:%v\n",i,j)
		}
	}
	fmt.Println("-----ok")

	/* 
		i:1, j:2
		i:1, j:3
		i:1, j:4
		i:3, j:2
		i:3, j:3
		i:3, j:4
		i:4, j:2
		i:4, j:3
		i:4, j:4
		i:5, j:2
		i:5, j:3
		i:5, j:4
		-----ok
	*/

}