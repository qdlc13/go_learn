package main
import(
	"fmt"
)
/*
继承(提高代码复用性拓展性)
多个结构体含有相同的字段和方法时，可以从这些结构体中抽象出结构体，在这个结构体定义相同的属性和方法，
其他结构体不需要重新定义这些属性和方法，只需要嵌套一个匿名结构体即可。（匿名结构体就是抽象出共有结构体的）

golang中如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现
继承特性。 
 */
 //定义动物(抽象共有的)
 type Animal struct{
	Age int
	Weight float32
 }
 //动物方法
 func(an *Animal) Shout(){
	fmt.Println("我可以叫")
 }
 func(an *Animal) ShowInfo(){
	fmt.Printf("我的年龄时%v我的体重%v\n",(*an).Age,(*an).Weight)
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

	(*cat).Animal.Age = 3//cat.Animal.Age也可以下同
	(*cat).Animal.Weight = 10.3
	(*cat).Animal.Shout()//我可以叫
	(*cat).Animal.ShowInfo()//我的年龄时3我的体重10.3
	(*cat).scratch()//我是小猫我可以抓人
 }