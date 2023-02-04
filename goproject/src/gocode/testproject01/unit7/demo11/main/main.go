package main
import(
	"fmt"
	"gocode/testproject01/unit7/demo11/model"
)
/* 
	封装
	封装就是把抽象出的字段和对字段的操作封装到一起，数据被保护在内部，
	程序的其他包只有通过被授权的操作方法才能对字段进行操作。

	封装的好处
	1.隐藏实现细节
	2.可以对数据进行验证保证安全合理

	golang实现封装
	1.建议将结构体，字段（属性）的首字母小写，其他包不能直接调用。
	2.给结构体所在的包提供一个工厂模式函数，首字母大写
	3.提供首字母大写的Get、Set方法用于对属性判断并赋值、取值

*/
func main() {
	//创建person结构体
	p := model.NewPerson("花花")
	p.SetAge(16)
	p.SetAge(200)//对不起你传入的年龄范围不正确

	/* fmt.Println((*p).name) 
	.\main.go:26:19: (*p).name undefined (type model.person has no field or method name)
	报错小写的属性不能直接访问
	*/
	fmt.Println(*p)				//{花花 16}
	fmt.Println((*p).GetAge())	//16


}