package main
import(
	"fmt"
)
func main(){
	//转义字符练习
	//\n换行
	fmt.Println("aaa\nbbb") 
	/* 
	aaa
	bbb
	*/
	//\b退格
	fmt.Println("aaa\bbbb")//aabbb
	
	//\r 光标回到本行开头 后续输入替换原有字符
	fmt.Println("aaa\rbbb")  //bbb

	//\t制表符 调到下一个tab位置
	fmt.Println("123456789")
	fmt.Println("aaaaa\tbbbbb")
    fmt.Println("aa\tbbbbb")
	/* 
	aaaaa   bbbbb
	aa      bbbbb
	*/

	//\"
	fmt.Println("\"Golang\"")//"Golang"






}