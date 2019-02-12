package main

import (
	"fmt"
)

func main() {
	// 数组的声明
	a:= [2]int{1,2}
	//a1 := [...]int{1,2,3}
	
	// 数组和切片的遍历
	for _,value := range a {
		fmt.Println(value)
	}
	
	
	// 切片的声明
	b:=[]int{1,2,3,4,5}
	
	
	// 数组和切片的比较
	fmt.Println("array:",a)
	fmt.Println(cap(a)) // 2 
	fmt.Println(len(a)) // 2
	
	fmt.Println("slice:",b)
	fmt.Println(cap(b)) //5
	fmt.Println(len(b)) //5
	
	c:= [10]int{1,2,3,4,5,6,7,8,9,0}
	d:=c[1:3]
	fmt.Println(d)
	fmt.Println(cap(d))
	fmt.Println(len(d))
	//fmt.Println(d[8]) //索引错了，索引讲究是长度
	fmt.Println(d[4:7]) // 切片没错，切片讲究的是容量
	

	//e := append(a,1) // 错误，append 只能应用于切片
	e :=append(d,3)
	fmt.Println(e) //2,3,3
	fmt.Println(c) // 原数组也改变了 [1 2 3 3 5 6 7 8 9 0]
 
	
}
