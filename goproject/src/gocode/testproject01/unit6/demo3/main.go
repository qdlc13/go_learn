package main
import(
	"fmt"
)
func main() {
	//定义map
	b := make(map[int]string)
	//增加
	b[20230130] = "张三"
	b[20230131] = "里斯"
	/* 
	【1】增加和更新操作:
	map["key"]= value  ——》 如果key还没有，就是增加，
	如果key存在就是修改。*/
	b[20230131] = "李四"
	b[20230132] = "王五"
	fmt.Println(b)//map[20230130:张三 20230131:李四 20230132:王五]

	/*
	【2】删除操作
	delete(map，"key") , delete是一个内置函数，如果key存在，
	就删除该key-value，如果k的y不存在，不操作，但是也不会报错*/	
	delete(b,20230132)
	delete(b,20230142)//删除没有的key也不会报错
	fmt.Println(b)

	/*
	【3】清空操作：
	（1）如果我们要删除map的所有key ,没有一个专门的方法一次删除，可以遍历一下key,逐个删除
	（2）或者map = make(...)，make一个新的，让原来的成为垃圾，被gc回收
	*/	
	b = make(map[int]string)//第二种清空方式
	fmt.Println(b)//map[]

	/*
	【4】查找操作：
	value ,bool := map[key]
	value为返回的value，bool为是否返回 ，要么true 要么false 
	*/
	value,flag := b[2222222]
	fmt.Println(value)//空的
	fmt.Println(flag)//false

	//遍历(长度获取len函数)
	b[20095452] = "张三"
    b[20095387] = "李四"
    b[20098833] = "王五"

	for k,v := range b {
		fmt.Printf("key为：%v value为%v \n",k,v)
	}
	/* 	
		key为：20095452 value为张三
		key为：20095387 value为李四
		key为：20098833 value为王五 
	*/
	fmt.Println("---------------------------")

	a := make(map[string]map[int]string) //map a的key是string类型value是map[int]string类型
	//赋值
	a["班级1"] = make(map[int]string,3)
	a["班级1"][20096677] = "露露"
    a["班级1"][20098833] = "丽丽"
    a["班级1"][20097722] = "菲菲"

	a["班级2"] = make(map[int]string)
	a["班级2"][20096677] = "露露"
    a["班级2"][20098833] = "丽丽"

	a["班级3"] = map[int]string{
		20095452 : "张三",
        20098765 : "李四",
	}

	//遍历
	for k1,v1:= range a {
		fmt.Println(k1)
		for k2,v2:= range v1{
				fmt.Printf("学生学号为：%v 学生姓名为%v \t",k2,v2)
		}
		fmt.Println()
	}

	/* 
		班级1
		学生学号为：20096677 学生姓名为露露     学生学号为：20098833 学生姓名为丽丽     学生学号为：20097722 学生姓名为菲菲
		班级2
		学生学号为：20096677 学生姓名为露露     学生学号为：20098833 学生姓名为丽丽
		班级3
		学生学号为：20095452 学生姓名为张三     学生学号为：20098765 学生姓名为李四
	*/

}