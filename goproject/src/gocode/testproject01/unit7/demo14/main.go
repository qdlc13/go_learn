package main
import(
	"fmt"
)

/* 	
	当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则，
	如果希望访问匿名结构体的字段和方法，可以通过匿名结构体的名字区分
*/

//定义动物(抽象共有的)
type Animal struct{
	age int
	Weight float32
 }
 //动物方法
 func(an *Animal) Shout(){
	fmt.Println("我可以叫")
 }
 func(an *Animal) showInfo(){
	fmt.Printf("我的年龄时%v我的体重%v\n",(*an).age,(*an).Weight)
 }

 //猫
 type Cat struct{
	//提高复用性，体现继承，加入匿名结构体->将Animal中的字段和方法都达到复用
	Animal
	age int
 }
 //cat特有方法
 func (c *Cat) scratch(){
	fmt.Println("我是小猫我可以抓人")
 }
 func(c *Cat) showInfo(){
	fmt.Printf("我是猫，我的年龄时%v我的体重%v\n",(*c).age,(*c).Weight)
 }
 


 func main(){
	//Cat结构体
	cat := &Cat{}
	cat.Weight = 5.7
	cat.age = 13 //就近原则

	cat.showInfo() //我是猫，我的年龄时13我的体重5.7  调用的就近cat的age和cat的showInfo函数
	cat.Animal.showInfo()//我的年龄时0我的体重5.7    调用的是Animal方法 age是0因为这是animal中的age

	cat.Animal.age = 20
	cat.Animal.showInfo()//我的年龄时20我的体重5.7
 }