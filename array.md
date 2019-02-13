# 数组
数组因为需要固定长度，因此在go里面使用较少，取而代之的是切片类型，除此数组的一些特性跟其他语言没什么差别，也是从 0 开始，每个数据绑定在一个下标的位置上
```go
a:= [3]int{1,2}
fmt.Println(a[0])
fmt.Println(len(a))


// 数组的遍历
for index,value := range a {
    fmt.Printf("%d : %d \n",index,value)
}
```
如果数组声明了长度但是没有赋值，其默认值为 0 
```go
a:= [3]int{1,2}
fmt.Println(a[2]) // 0
```
另外的一种数组声明方式如下所示：
```go
// 另外的一种数组声明
a1 := [...]int{1,2,3}
fmt.Printf("%T \n",a1) //[3]int 
```
一般的数组以数字作为其下标，在go中，我们也可以指定一些更加语义化的变量作为其下标
```go
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
```
有时候，我们声明数组，不想从第一个位置开始赋值，而是需要从中间开始赋值
```go
//从数组的某一个位置开始赋值
a2 := [...]int{3:1,2,3,4,5,6} //[0 0 0 1 2 3 4 5 6]
fmt.Println(a2)
```
如果数组的元素是可以比较的，那么数组就可以直接进行比较，但是前提是数组的长度必须要保持一致，否则就会出错
```go
// 数组的比较
c1 := [2]int {1,2}
c2 := [...]int {1,2}
c3 := [2]int {1,3}
c4 := [3]int {1,2}

fmt.Println(c1 == c2) // true
fmt.Println(c1 == c3) // false
fmt.Println(c1 == c4) // invalid operation: c1 == c4 (mismatched types [2]int and [3]int)
```
对于一些似是而非的数组，比如 Sum256 函数，返回的是一个 [32]int 的数组，但是却是一串数字
```go
d1 := sha256.Sum256([]byte("x"))
d2 := sha256.Sum256([]byte("X"))
fmt.Printf("%x\n%x\n%t\n%T",d1,d2,d1==d2,d1)
fmt.Println(d1[1]) // 113
/**
2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
false
[32]uint8%
 */
```
当数组做为一个函数的参数的时候，其实是经过了一层的拷贝。函数改变的数组的值，对于原数组没有任何的影响。为了性能的考虑，这个参数最好为数组的指针，这样就避免了拷贝的发生
```go
func change(a *[3]int){
	a[1] = 1
	a[2] = 2
}

e1 := [3]int {4,5,6}
change(&e1)
fmt.Println(e1) // [4 1 2] 改变了原数组
```
## Exercise 4.1 
```go
// 4.1

func countNumberOfHash(d *[32]uint8) int {
	fmt.Println(*d) //[36 129 166 60 133 166 44 248 137 210 177 73 241 165 46 152 90 147 65 117 1 115 254 1 239 245 12 194 123 89 65 181]
	return len(*d) //32
}

d1 := sha256.Sum256([]byte("xxxx"))
count := countNumberOfHash(&d1)
fmt.Println(count)  //32
```

## Exercise 4.2 
```go


```