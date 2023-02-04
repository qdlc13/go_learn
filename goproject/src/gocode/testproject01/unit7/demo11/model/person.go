package model
import(
	"fmt"
)
type person struct{
	name string
	age int //不能直接访问

}
//定义工厂方式的函数，相当于构造器
func NewPerson(name string) *person{
	return &person{
		name : name,
	}
}
//定义set和get方法，对age字段进行封装
//因为在方法中可以加一系列限制操作确保被封装字段安全合理性
func (p *person) SetAge(age int){
	if age > 0 && age < 150 {
		(*p).age = age
	}else {
		fmt.Println("对不起你传入的年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return (*p).age
}