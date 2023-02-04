package main
import(
	"fmt"
)
func main(){
//（键值循环） for range结构是Go语言特有的一种的迭代结构，在许多情况下都非常有用，
//for range 可以遍历数组、切片、字符串、map 及通道，for range 语法上类似于其它语言中的 
//foreach 语句，一般形式为：
// for key, val := range coll {
//     ...
// }
	var str string = "hello golang你好"
	//for循环 按照字节遍历中文会拆成三个字节
	for i := 0 ; i < len(str) ; i++ {
		fmt.Printf("%c\n",str[i])    //中文部分乱码
	}
	/* 
	h
	e
	l
	l
	o

	g
	o
	l
	a
	n
	g
	ä
	½
	
	å
	¥
	½ 
	*/

	fmt.Println("==================================================")

	//for range
	for i , value := range str {
		fmt.Printf("索引为：%d,具体值为: %c \n",i,value)
	}
	/* 
	索引为：0,具体值为: h
	索引为：1,具体值为: e
	索引为：2,具体值为: l
	索引为：3,具体值为: l
	索引为：4,具体值为: o
	索引为：5,具体值为:
	索引为：6,具体值为: g
	索引为：7,具体值为: o
	索引为：8,具体值为: l
	索引为：9,具体值为: a
	索引为：10,具体值为: n
	索引为：11,具体值为: g
	索引为：12,具体值为: 你
	索引为：15,具体值为: 好  注意是15
 	*/
	//对str进行遍历，遍历的每个结果的索引值被i接收，每个结果的具体数值被value接收
    //遍历对字符进行遍历的
}