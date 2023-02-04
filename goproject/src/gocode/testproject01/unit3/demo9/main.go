package main
import(
	"fmt"
)
func main(){
	var sum int = 0
	for i := 1 ; i <= 100 ; i++ {
		sum += i
		fmt.Println(sum)
		if sum >= 300 {
			break
		}
	}
	//1.switch分支中，每个case分支后都用break结束当前分支，但是在go语言中break可以省略不写。
	//2.break可以结束正在执行的循环
	//break的作用结束离它最近的循环
	for i := 1 ; i <= 5 ; i++ {
		for j := 2 ; j <= 4 ; j++ {
			fmt.Printf("i:%v,j:%v\n",i,j)
			if i == 2 && j ==2 {
				break
			}
		}
	}

	/* 
		1
		3
		6
		10
		15
		21
		28
		36
		45
		55
		66
		78
		91
		105
		120
		136
		153
		171
		190
		210
		231
		253
		276
		300
		i:1,j:2
		i:1,j:3
		i:1,j:4
		i:2,j:2
		i:3,j:2
		i:3,j:3
		i:3,j:4
		i:4,j:2
		i:4,j:3
		i:4,j:4
		i:5,j:2
	*/

}