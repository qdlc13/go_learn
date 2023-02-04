package main
import(
	"fmt"
	"errors"
)
func main() {
	//自定义错误：需要调用errors包下的New函数：函数返回error类型
	/* 
		有一种情况：程序出现错误以后，后续代码就没有必要执行，
		想让程序中断，退出程序：
		借助：builtin包下内置函数：panic
	*/
	err := test()
	if err != nil {
		fmt.Println("自定义错误：",err)
		panic(err)
	}
	fmt.Println("上面除法正常执行")//未执行
	fmt.Println("执行新的逻辑")//
}

func test() (err error){
	num1 := 10
	num2 := 0
	if num2 == 0 {
		//抛出自定义错误
		//errors包下
		return errors.New("除数不能为0哦")
	}else {//除数不为0正常执行
		result := num1 / num2
		fmt.Println(result)
		//没有错误返回零值
		return nil
	}
}
/* 
自定义错误： 除数不能为0哦
panic: 除数不能为0哦

goroutine 1 [running]:
main.main()
        C:/Users/qdliu/Desktop/goproject/src/gocode/testproject01/unit4/demo20/main.go:16 +0x128
exit status 2
*/