# Go语言内置容器
Go语言的内置容器主要有数组、切片、映射。
## 数组
数组是具有相同类型且长度固定的一组数据项序列，这组数据项序列对应存放在内存中的一块连续区域中。
数组中存放的元素类型可以是整型、字符型或其他自定义类型。数组在使用前需声明，声明中必须指定数组的大小且数组大小之后是不可变的。
数组元素可以通过数组下标来读取或者修改。
### 声明数组
数组声明格式如下：
```go
var 数组变量名 [数组长度] 元素类型
```

### 初始化数组
数组可以在声明时进行赋值，举例：
```go
var student = [3] string {"Tom" , "Ben" , "Peter"}
```
注意：使用这种方式进行初始化数组，需要保证大括号里面的元素数量和数组大小保持一致。
还可以通过以下方式进行初始化操作：
```go
var student = [...] string {"Tom", "Ben" , "Peter"}
```
### range 关键字
range关键字的主要作用是配合for关键字对数组以及切片、映射等数据结构进行迭代。
举例：
```go
package main

import "fmt"

func main() {
    var num = [...] int {1, 2, 3, 4}
    for k, v := range num {
        fmt.Println("变量k：",k,"变量v：",v)
    }
}
```
执行结果：
```go
变量k: 0 变量v: 1
变量k: 1 变量v: 2
变量k: 2 变量v: 3
```
|range表达式|第一返回值|第二返回值|
|-----------|-------|-----------|
|数组|元素下标|元素值|
|切片|元素下标|元素值|
|映射|键|值|
|通道|元素|N/A|

## 切片
相对于数组，切片是一种更方便和强大的数据结构，它同样表示多个同类型的元素的连续集合，但是切片本身并不存储任何元素，而只是对现有的数组的引用。
切片的结构包括：地址、长度、容量。
1.地址：切片的地址一般是指切片中第一个元素所指向的内存地址，用十六进制表示。
2.长度：切片中实际存在元素的个数。
3.容量：切片的起始元素开始到其底层数组中的最后一个元素的个数。
切片的长度和容量都是不固定的，可以通过追加元素使切片的长度和容量增大。
### 从数组生成一个新的切片
从数组或切片生成新的切片的语法格式：
```go
slice [开始位置：结束位置]
```
使用len（）函数和cap（）函数可以获取切片的长度和容量。

### 直接生成一个切片
1.声明切片
```go
var 切片变量名 [] 元素类型
```
举例：
```go
package main

import "fmt"

func main() {
    var student [] int
    fmt.Println("student切片：",student)
    fmt.Println("student切片长度：",len(student))
    fmr.Println("student切片容量：",cap(student))
    fmt.Prinfln("判定student切片是否为空：",student==nil)
}
```

结果：
```go
student切片： []
student切片长度： 0
student切片容量:  0
判定student是否为空： true
```
2.初始化切片
（1）、在声明时初始化
（2）、使用make（）函数初始化
```go
make([]元素类型,切片长度,切片容量)
```
注意：切片的容量值必须要大于等于切片长度值。
### 为切片添加元素
使用append（）函数来对切片元素进行添加。当切片不能再容纳其他元素时，即当前切片长度值等于切片容量值，下一次再使用append（）函数对切片进行添加时，容量会按2倍数进行扩充。
举例：
```go
package main

import "fmt"

func main() {
    student := make([]int,1,1)
    for i := 0; i < 8; i++ {
        student = append(student,i)
        fmt.Println("当前切片长度：",len(student),"当前切片容量：",cap(student))
    }
}
```
结果：
```go
当前切片长度： 2 当前切片容量：2
当前切片长度： 3 当前切片容量：4
当前切片长度： 4 当前切片容量：4
当前切片长度： 5 当前切片容量：8
当前切片长度： 6 当前切片容量：8
当前切片长度： 7 当前切片容量：8
当前切片长度： 8 当前切片容量：8
当前切片长度： 9 当前切片容量：16
```



### 从切片删除元素
由于Go语言没有删除切片元素提供方法，所以需要我们手动将删除点前后的元素连接起来，从而实现对切片元素的删除。
举例：
```go
package main

import "fmt"

func main() {
    var student = []string{"Tom","Ben","Peter","Danny"}
    student = append(student[0:1],student[2:]...)
    fmt.Println("student切片：",student)
    fmt.Println("student切片长度：",len(student))
    fmr.Println("student切片容量：",cap(student))
}
```
结果：
```go
student切片： [Tom Peter Danny]
student切片长度： 3
student切片容量： 4
```
其中append()函数传入的省略号代表按student切片展开。
如果需要清空切片中的元素，可以把切片的开始下标和结束下标设为0.
### 遍历切片
与数组遍历类似。
## 映射
映射是一种无序的键值对集合，map的键类似于索引，指向数据的值。
### 声明映射
map的声明格式如下：
```go
var map [键类型] [值类型]
```

### 初始化映射
(1)在声明的同时初始化
举例：
```go
package main

import "fmt"

func main() {
    var studentScoreMap = map [string] [int] {
        "Tom" : 80,
        "Ben" : 90,
        "Peter" : 100,
    }
    fmt.Println("studentScoreMap")
}
```
结果：
```go
map [Tom：80 ben: 90 Peter: 100]
```
(2)使用make（）函数初始化
格式如下：
```go
make (map[键类型] 值类型, map容量)
```
### 遍历映射
和遍历数组一样。
### 从映射中删除键值对
Go语言通过delete（）函数来对map中指定键值进行删除操作，delete（）函数格式如下：
```go
delete(map, 键)
```
举例：
```go
package main

import "fmt"

func main() {
    var studentScoreMap map[string] int
    studentScoreMap = make (map [string] int)
    studentScoreMao ["Tom"]=80
    studentScoreMap ["Ben"]=90
    studentScoreMap ["Peter"]=100
    delete(studentScoreMap,"Tom")
    fmt.Println("studentScoreMap")
}
```
结果：
```go
map [Ben:80 Peter:100]
```