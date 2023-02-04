package main
import(
	"fmt"
)
/* 
	结构体可以使用匿名结构体的所有字段和方法 即首字母大写或者小写的字段方法都可以使用

	匿名结构体字段访问可以简化

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
 }
 //cat特有方法
 func (c *Cat) scratch(){
	fmt.Println("我是小猫我可以抓人")
 }
 


 func main(){
	//Cat结构体
	cat := &Cat{}

	//cat.age->系统先在cat对应结构体寻找是否有age字段，没有就去嵌入的结构体类型中找age
	//匿名结构体字段访问可以简化
	/* 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则，
	如果希望访问匿名结构体的字段和方法，可以通过匿名结构体的名字区分 */
	(*cat).age = 3//cat.Animal.Age也可以下同   //等价(*cat).Animal.age = 3
	(*cat).Weight = 10.3                      //等价(*cat).Animal.Weight = 10.3
	(*cat).Shout()//我可以叫                   //等价(*cat).Animal.Shout()
	(*cat).showInfo()//我的年龄时3我的体重10.3  //等价(*cat).Animal.showInfo()
	(*cat).scratch()//我是小猫我可以抓人
 }