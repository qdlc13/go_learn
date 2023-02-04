package main
import(
	"fmt"
)
func main(){
	var num1 int = 10
	fmt.Println(num1)//10
	var num2 int = (10 + 20) % 3 + 3 - 7
	fmt.Println(num2)//-4

	var num3 int = 10
	num3 += 20
    fmt.Println(num3)//30

	//交换数值
	var a int = 8
	var b int = 4
	fmt.Printf("a = %v,b = %v\n",a,b)//8 4

	var t int
	t = a
	a = b
	b = t
	fmt.Printf("a = %v,b = %v\n",a,b)//4 8

	//关系运算
	fmt.Println(5==9)//false
	fmt.Println(5!=9)//true
	fmt.Println(5>9)//false
	fmt.Println(5<9)//true
	fmt.Println(5>=9)//false
	fmt.Println(5<=9)//true
	fmt.Println("=======================================================")
	//逻辑运算
	//与逻辑：&& :两个数值/表达式只要有一侧是false，结果一定为false
    //也叫短路与：只要第一个数值/表达式的结果是false，
	//那么后面的表达式等就不用运算了，直接结果就是false  -->提高运算效率
	fmt.Println(true&&true)//true
    fmt.Println(true&&false)//false
    fmt.Println(false&&true)//false
    fmt.Println(false&&false)//false
	fmt.Println("=======================================================")
	//或逻辑：||:两个数值/表达式只要有一侧是true，结果一定为true
    //也叫短路或：只要第一个数值/表达式的结果是true，
	//那么后面的表达式等就不用运算了，直接结果就是true -->提高运算效率
	fmt.Println(true||true)//true
    fmt.Println(true||false)//true
    fmt.Println(false||true)//true
    fmt.Println(false||false)//false
	fmt.Println("=======================================================")
	//非逻辑：取相反的结果：
	fmt.Println(!true)//false
	fmt.Println(!false)//true

}