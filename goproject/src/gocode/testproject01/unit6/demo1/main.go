package main
import(
	"fmt"
)
func main() {
	//映射（map）, Go语言中内置的一种类型，它将键值对相关联，
	//我们可以通过键 key来获取对应的值 value。 类似其它语言的集合

	/* 
		基本语法
	var map变量名 map[keytype]valuetype
	PS：key、value的类型：bool、数字、string、指针、channel 、
	还可以是只包含前面几个类型的接口、结构体、数组
	PS：key通常为int 、string类型，value通常为数字（整数、浮点数）、string、map、结构体
	PS：key：slice、map、function不可以

	代码：   
	map的特点：
	（1）map集合在使用前一定要make
	（2）map的key-value是无序的
	（3）key是不可以重复的，如果遇到重复，后一个value会替换前一个value
	（4）value可以重复的
	*/

	//定义map变量
	var a map[int]string
	//只声明map内存是没有分配空间
    //必须通过make函数进行初始化，才会分配空间：
	a = make(map[int]string,10)
	//将键值对存入map中：
	a[20095452] = "张三"
    a[20095387] = "李四"
    a[20097291] = "王五"
    a[20095387] = "朱六"
    a[20096699] = "张三"

	//输出集合
	fmt.Println(a)//map[20095387:朱六 20095452:张三 20096699:张三 20097291:王五]

}