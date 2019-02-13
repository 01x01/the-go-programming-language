package code

import (
	"crypto/sha256"
	"fmt"
)

func change(a *[3]int){
	a[1] = 1
	a[2] = 2
}


func Arrays() {
	// 数组的声明
	a:= [3]int{1,2}
	fmt.Println(a[0])
	fmt.Println(len(a))


	// 数组的遍历
	for index,value := range a {
		fmt.Printf("%d : %d \n",index,value)
	}

	fmt.Println(a[2]) // 0

	// 另外的一种数组声明
	a1 := [...]int{1,2,3}
	fmt.Printf("%T \n",a1) //[3]int

	// 语义化下标
	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string {
		USD: "$",
		EUR: "%",
		GBP: "&",
		RMB: "¥",
	}
	fmt.Println(symbol[RMB]) // ¥
	//从数组的某一个位置开始赋值
	a2 := [...]int{3:1,2,3,4,5,6} //[0 0 0 1 2 3 4 5 6]
	fmt.Println(a2)

	// 数组的比较
	c1 := [2]int {1,2}
	c2 := [...]int {1,2}
	c3 := [2]int {1,3}
	//c4 := [3]int {1,2}

	fmt.Println(c1 == c2) // true
	fmt.Println(c1 == c3) // false
	//fmt.Println(c1 == c4) // invalid operation: c1 == c4 (mismatched types [2]int and [3]int)


	d1 := sha256.Sum256([]byte("xxxx"))
	d2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n",d1,d2,d1==d2,d1)
	fmt.Println(d1[1])
	/**
	2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	false
	[32]uint8%
	 */

	 e1 := [3]int {4,5,6}
	 change(&e1)
	 fmt.Println(e1) // [4 1 2] 改变了原数组

	 //4.1
	count := countNumberOfHash(&d1)
	fmt.Println(count)

}

// 4.1
func countNumberOfHash(d *[32]uint8) int {
	fmt.Println(*d) //[36 129 166 60 133 166 44 248 137 210 177 73 241 165 46 152 90 147 65 117 1 115 254 1 239 245 12 194 123 89 65 181]
	return len(*d) //32
}
//4.2
func asha256(s string)[32]uint8{
	return sha256.Sum256([]byte(s))
}
