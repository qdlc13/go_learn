package main
import(
	"fmt"
)
/* 
	Golang中没有专门的字符类型，如果要存储单个字符(字母)，一般使用byte来保存。
	Golang中字符使用UTF-8编码
 */
func main(){
	//输出字符十进制ASCII码
	var c1 byte = 'a'
	fmt.Println(c1)  //97
	var c2 byte = '6'
	fmt.Println(c2)  //54
	var c3 byte = '('
	fmt.Println(c3 + 20) //40 + 20 = 60

	/* 
	汉字字符，底层对应的是Unicode码值
    对应的码值为20320，byte类型溢出，能存储的范围：可以用int
    总结：Golang的字符对应的使用的是UTF-8编码
	（Unicode是对应的字符集，UTF-8是Unicode的其中的一种编码方案） 
	
	var c4 byte = '你'
	# command-line-arguments
	.\main.go:23:16: cannot use '你' (untyped rune constant 20320) 
	as byte value in variable declaration (overflows)
	
	*/
	var c4 int = '你'
	fmt.Println(c4)//20320

	var c5 byte = 'S'
	fmt.Printf("c5具体字符是：%c\n",c5)//c5具体字符是：S

	var c6 int = '山'
	fmt.Printf("c6具体字符是：%c\u000a",c6)//c6具体字符是：山

	var c7,c8 int = 'f','山'
	fmt.Printf("c7,c8具体字符是：%c %c",c7,c8)//c7,c8具体字符是：f 山
}