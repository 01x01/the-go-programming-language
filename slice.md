切片写作 []T 其中T为类型，看起来像是一个数组，但是没有指定长度。
数组和切片是密切联系的，切片相当于是数组的一个子集，一个切片有三种属性，指针，长度，容量
长度是切片的个数
容量是切片后，数组剩下的个数。一个简单的例子：
```go
func Slice(){
	fmt.Println("Slice....\n")
	arrs := [5]string{"Hello","World","nihao","aliyadou","xixix"}
	slices := arrs[2:4]
	fmt.Printf("length: %d\n",len(slices))
	fmt.Printf("cap: %d\n",cap(slices))

}
```
其输出为
```go
Slice....
length: 2
cap: 3
```
切片的长度为 2 ，容量为3 说明容量的计算，是从数组索引为 2 的位置开始算起。

