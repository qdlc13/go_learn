package main
import(
	"fmt"
)
func exNum(num1 int){
	fmt.Println("hha")
}
func exNum(num1 int,num2 int){
	fmt.Println("hha")
}
func main(){
	//exNum(2) go不支持重载
	/* .\main.go:8:6: exNum redeclared in this block    
    .\main.go:5:6: other declaration of exNum */

}