// 指针
package code

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	p := 12
	fmt.Println(p)
	fmt.Println(&p)
	
	// 声明指针
	var p2 *int = &p 
	fmt.Println(p2)
	fmt.Println(*p2) // 只有指针类型才可以使用 * 
}
