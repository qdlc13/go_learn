package main
import(
	"fmt"
	"strconv"
	"strings"
)
func main() {
	//字符串
	str := "golang 你好"
	fmt.Println(len(str))//13字节 汉字是utf-8字符集占三个字节

	//字符串遍历
	//键值循环for-range
	for i , value := range str {
		fmt.Printf("索引为%d,具体的值为%c \n",i,value)
	}
	/* 
		索引为0,具体的值为g 
		索引为1,具体的值为o
		索引为2,具体的值为l
		索引为3,具体的值为a
		索引为4,具体的值为n
		索引为5,具体的值为g
		索引为6,具体的值为
		索引为7,具体的值为你
		索引为10,具体的值为好
	*/

	//r:=[]rune(str) 字符串转切片 然后遍历
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n",r[i])
	} 
	/* 
		g
		o
		l
		a
		n
		g

		你
		好
	*/

	//字符串转整数num, err := strconv.Atoi("66") 
	num, _ := strconv.Atoi("66") 
	fmt.Println(num) //66

	//整数转字符串
	str2 := strconv.Itoa(6887)
	fmt.Println(str2)

	//查找子串是否在指定的字符串中
	fmt.Println(strings.Contains("javaandgolang", "go"))//true

	//统计一个字符串有几个指定的子串
	count2 := strings.Count("javaandgolang","a") 
	fmt.Println(count2)//4

	//不区分大小写的字符串比较
	flag := strings.EqualFold("go" , "Go")
	fmt.Println(flag)//true

	//区分大小写进行字符串比较
	fmt.Println("hello" == "Hello")//false

	//返回子串在字符串第一次出现的索引值，如果没有返回-1 :
	fmt.Println(strings.Index("javaandgolang" , "a"))//1
	fmt.Println(strings.Index("javaandgolang" , "np"))//-1

	//字符串替换 strings.Replace("goandjavagogo", "go", "golang", n)
	//n可以指定你希望替换几个,如果n=-1表示全部替换，替换两个n就是2
	s1 := strings.Replace("goandjavagogo", "go", "golang", -1)
	s2 := strings.Replace("goandjavagogo", "go", "golang", 2)
	fmt.Println(s1)//golangandjavagolanggolang
	fmt.Println(s2)//golangandjavagolanggo

	//按照指定的某个字符，为分割标识，将一个学符串拆分成字符串数组:
	s3 := strings.Split("go-python-java", "-")
	fmt.Println(s3)//[go python java]

	//将字符串的字母进行大小写的转换:
	fmt.Println(strings.ToLower("GOo"))// goo 
	fmt.Println(strings.ToUpper("goO"))//GOO

	//将字符串左右两边的空格去掉:
	fmt.Println(strings.TrimSpace("  go and java    "))//go and java

	//将字符串左右两边指定的字符去掉: 遇到空格不停
	fmt.Println(strings.Trim("~    ~yy~~~golang~~yy~~~  ~~~   ", " ~"))
	//yy~~~golang~~yy

	//将字符串左边指定的字符去掉: 遇到空格停止 
	fmt.Println(strings.TrimLeft("~ ~ ~yy~~~golang~~yy~~~  ~~~   ", "~"))
	// ~ ~yy~~~golang~~yy~~~  ~~~

	//将字符串右边指定的字符去掉: 遇到空格停止
	fmt.Println(strings.TrimRight("~    ~yy~~~golang~~yy~~~  ~~~", "~"))
	//~    ~yy~~~golang~~yy~~~

	//判断字符串是否以指定的字符串开头: 
	fmt.Println(strings.HasPrefix("http://java.sun.com/jsp/jstl/fmt", "http"))//true
	fmt.Println(strings.HasPrefix("http://java.sun.com/jsp/jstl/fmt", "htwp"))//false

	//判断字符串是否以指定的字符串结束: 
	fmt.Println(strings.HasSuffix("demo.png", ".png"))//true
	fmt.Println(strings.HasSuffix("demo.png", ".dpng"))//false
}