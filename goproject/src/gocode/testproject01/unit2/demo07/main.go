package main
import(
	"fmt"
)
func main(){
	//字符串
	var s1 string = "你好啊啊啊啊啊"
	fmt.Println(s1)		//你好啊啊啊啊啊
	fmt.Printf("%s\n",s1) //你好啊啊啊啊啊
	fmt.Println(s1[0]) //228 应该也是乱码
	fmt.Printf("%c\n",s1[0]) //乱码ä

	var s0 string = "abcdef"
	fmt.Println(s0)		//abcdef
	fmt.Printf("%s\n",s0) //abcdef
	fmt.Println(s0[1]) //98
	fmt.Printf("%c\n",s0[1]) //b

	//字符串不可变 定义好值不能改变
	//https://blog.csdn.net/qq_27681741/article/details/127563659
	var s2 string = "abc"
	//s2 = "defwwwwww"  //defwwwwww 可以

    /* s2[0] = "r"     //报错
	# command-line-arguments
	.\main.go:22:5: cannot assign to 
	s2[0] (value of type byte) */
	fmt.Println(s2) 
	

	//3.字符串的表示形式：
    //（1）如果字符串中没有特殊字符，字符串的表示形式用双引号
    //var s3 string = "asdfasdfasdf"
    //（2）如果字符串中有特殊字符，字符串的表示形式用反引号 ``  格式原样输出

	var s4 string = `
        package main
        import "fmt"
        
        func main(){
                //测试布尔类型的数值：
                var flag01 bool = true
                fmt.Println(flag01)
        
                var flag02 bool = false
                fmt.Println(flag02)
        
                var flag03 bool = 5 < 9
                fmt.Println(flag03)
        }
        `
        fmt.Println(s4)

		//字符串的拼接
		var s5 string = "dss" + "333"
		s5 += "fff"
		fmt.Println(s5) //dss333fff

		/*  
		当一个字符串过长的时候：注意：+保留在上一行的最后 否则会报错 .\main.go:61:5: invalid 
		operation: operator + not defined on "def" (untyped string constant)
		*/
        var s6 string = "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + 
		"def"+ "abc" + "def" + "abc" + "def"+ "abc" + "def" + "abc" + "def"+
        "abc" + "def" + "abc" + "def"+ "abc" + "def" + "abc" + "def"+ "abc" +
        "def" + "abc" + "def"+ "abc" + "def" + "abc" + "def"+ "abc" + "def" + 
        "abc" + "def"+ "abc" + "def"
		//abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdef
        fmt.Println(s6)
}