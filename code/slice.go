package code

import "fmt"

func Slice(){
	fmt.Println("Slice....\n")
	arrs := [5]string{"Hello","World","nihao","aliyadou","xixix"}
	slices := arrs[2:4]
	fmt.Printf("length: %d\n",len(slices))
	fmt.Printf("cap: %d\n",cap(slices))
	months := [13]string{
		1:"January",
		2:"Febuary",
		3:"March",
		4:"April",
		5:"May",
		6:"June",
		7:"July",
		8:"August",
		9:"September",
		10:"October",
		11:"November",
		12:"December",
	}

	// 声明切片
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2) //[April May June]
	fmt.Println(summer) // [June July August]

	//循环切片
	for _,value := range summer{
		for _,value2 := range Q2{
			if value == value2{
				fmt.Printf("%s in Q2 and summer\n",value) //June in Q2 and summer
			}
		}
	}
	// 切片超过容量会出错，但是超过长度却是正常的
	//fmt.Println(summer[:20]) // 出错
	fmt.Println(summer[:5]) //  summer 的长度为 3，但是却可以延伸到 5

	// 切片包含的是数组的指针
	a := [10]int {1,2,3,4,5,6,7,8,9,10}
	reverse(a[:])
	fmt.Println(a)

	// 不像数组，切片是不可以比较的，因此我们要自己准备比较函数
	b1 := []int {1,2,3,4,5}
	d1 := []int {7,8,9,10,11}
	//fmt.Println(b1 == d1) //error
	fmt.Println(Euqal(b1,d1))

	// 因为切片的底层是数组，而数组随时可能经过变化，所以最好的方式是禁止切片的比较
	// 唯一正常的比较是跟nil的比较，一个空的切片，就是一个nil
	var anil []int
	fmt.Println( anil == nil) // true


}
func Euqal(a,b []int)bool{
	if len(a) != len(b) {
		return false
	}

	for i := range a{
		if a[i] != b[i]{
			return false
		}
	}
	return true
}
func reverse(a []int){
	for i,j:=0,len(a)-1;i<j;i,j = i+1,j-1{
		a[i],a[j] = a[j],a[i]
	}
}

