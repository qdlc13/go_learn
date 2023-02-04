package main
import (
	"fmt"
	"time"
)
func main() {
	//时间和日期的函数，需要到入time包，所以你获取当前时间，就要调用函数Now函数
	now := time.Now()
	//Now()返回值是一个结构体，类型是：time.Time
	fmt.Printf("%v ...对应的类型是：%T\n",now,now)
	//2023-01-25 20:20:20.9458276 +0800 CST m=+0.002732201 ...对应的类型是：time.Time
	
	//调用结构体中的方法：
	fmt.Printf("年:%v \n",now.Year())//年:2023 
	fmt.Printf("月:%v \n",now.Month())//月:January
	fmt.Printf("月:%v \n",int(now.Month()))//月:1
	fmt.Printf("日:%v \n",now.Day())//日:25
	fmt.Printf("时:%v \n",now.Hour())
	fmt.Printf("分:%v \n",now.Minute())
    fmt.Printf("秒:%v \n",now.Second())

	//Printf将字符串直接输出：
	fmt.Printf("当前年月日： %d-%d-%d 时分秒：%d:%d:%d  \n",now.Year(),now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())//当前年月日： 2023-1-26 时分秒：15:48:30

	//Sprintf可以得到这个字符串，以便后续使用：
	datestr := fmt.Sprintf("当前年月日： %d-%d-%d 时分秒：%d:%d:%d  \n",now.Year(),now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())

	fmt.Printf(datestr)//当前年月日： 2023-1-26 时分秒：15:48:30


	//按照指定格式：
	datestr2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(datestr2)//2023/01/26 16/10/48

	//选择任意的组合都是可以的，根据需求自己选择就可以（自己任意组合）。
	datestr3 := now.Format("2006 15:04")
	fmt.Println(datestr3)//2023 16:11




}