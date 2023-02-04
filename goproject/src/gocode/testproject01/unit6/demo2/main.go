package main
import(
	"fmt"
)
func main() {
	//创建map方式
	var a map[int]string
	//只声明map内存是没有分配空间
    //必须通过make函数进行初始化，才会分配空间：
	a = make(map[int]string,10) //map可以存放10个键值对
	//将键值对存入map中：
	a[20095452] = "张三"
	a[20095387] = "李四"
	//输出集合
	fmt.Println(a)//map[20095387:李四 20095452:张三]

	//方式二
	b := make(map[int]string)
	b[20095452] = "张三"
    b[20095387] = "李四"
    fmt.Println(b)//map[20095387:李四 20095452:张三]

	//方式三
	c := map[int]string{
		20095452 : "张三",
        20098765 : "李四",
	}
	c[20095387] = "王五"
    fmt.Println(c)//map[20095387:王五 20095452:张三 20098765:李四]


}