1 数组介绍
数组是同一类型元素的集合。例如，整数集合 5,8,9,79,76 形成一个数组。Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。（注：当然，如果是 interface{} 类型数组，可以包含任意类型） 。

2 数组常见操作
一个数组的表示形式为 [n]T。n 表示数组中元素的数量，T 代表每个元素的类型。元素的数量 n 也是该类型的一部分 。

2.1 数组初始化
一维数组初始化如下

func main() {
    var a [4]int    //元素自动初始化为零[0 0 0 0]
    b := [4]int{2, 5}  //未提供初始化值得元素自动初始化为0  [2 5 0 0]
    c := [4]int{5, 3: 10} //可指定索引位置初始化 [5 0 0 10];  3: 10 --> 数组下标为3的值设置为10
    d := [...]int{1, 2, 3} //编译器按初始化值数量确定数组长度 [1 2 3]
    e := [...]int{10, 3: 100} //支持索引初始化，但注意数组长度与此有关 [10 0 0 100]
    fmt.Println(a, b, c, d, e)
}

对于结构等复合类型，可省略元素初始化类型标签

package main

import "fmt"

func main() {
    type user struct {
        name string
        age  byte
    }

    d := [...]user{
        {"tom", 20},// 可省略元素类型。
        {"lee", 18},// 别忘了最后一行的逗号。
    }

    fmt.Printf("%#v\n", d)
}
/*output
[2]main.user{main.user{name:"tom", age:0x14}, main.user{name:"lee", age:0x12}}
*/

在定义多维数组时，仅第一维度允许使用“…”

package main

import "fmt"

func main() {
    a := [2][2]int{
        {1, 2},
        {3, 4},
    }

    b := [...][2]int{
        {10, 20},
        {30, 40},
    }

    c := [...][2][2]int{   //三维数组
        {
            {1, 2},
            {3, 4},
        },
        {
            {10, 20},
            {30, 40},
        },
    }

    fmt.Println(a)  //[[1 2] [3 4]]
    fmt.Println(b)  //[[10 20] [30 40]]
    fmt.Println(c)  //[[[1 2] [3 4]] [[10 20] [30 40]]]
}

多维数组定义

package main

import (
    "fmt"
)

var arr0 [5][3]int
var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func main() {
    a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
    b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
    fmt.Println(arr0, arr1)
    fmt.Println(a, b)
}

/*
output
[[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]] [[1 2 3] [7 8 9]]
[[1 2 3] [4 5 6]] [[1 1] [2 2] [3 3]]
 */

2.2 数组索引
数组的索引从 0 开始到 length - 1 结束

func main() {
    var a [3]int //int array with length 3
    a[0] = 12    // array index starts at 0
    a[1] = 78
    a[2] = 50
    fmt.Println(a)
}

2.3 数组是值类型
Go 中的数组是值类型而不是引用类型。这意味着当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本。如果对新变量进行更改，则不会影响原始数组。

func main() {
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)  //a is  [USA China India Germany France]
    fmt.Println("b is ", b)  //b is  [Singapore China India Germany France]
}

上述程序中，a 的副本被赋给 b。在第 4 行中，b 的第一个元素改为 Singapore。这不会在原始数组 a 中反映出来。

同样，当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。

package main

import "fmt"

func changeLocal(num [5]int) {
    num[0] = 55
    fmt.Println("inside function ", num)
}

func main() {
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num) 
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
}

/*output
before passing to function  [5 6 7 8 8]
inside function  [55 6 7 8 8]
after passing to function  [5 6 7 8 8]
*/

在上述程序的 13 行中, 数组 num 实际上是通过值传递给函数 changeLocal，数组不会因为函数调用而改变。

值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。

package main

import (
    "fmt"
)

func test(x [2]int) {
    fmt.Printf("x: %p\n", &x)
    x[1] = 1000
}
func main() {
    a := [2]int{}
    fmt.Printf("a: %p\n", &a)
    test(a)
    fmt.Println(a)
}
/*
output:
a: 0xc042062080
x: 0xc0420620c0
[0 0]
 */

2.4 数组长度和元素数量
通过将数组作为参数传递给 len 函数，可以得到数组的长度。 cap可以得到元素数量

package main

import "fmt"

func main() {
    a := [...]float64{67.7, 89.8, 21, 78}
    fmt.Println("length of a is", len(a)) //length of a is 4
    fmt.Println("num of a is",cap(a)) //num of a is 4
}

注意：内置函数len和cap都返回第一维度长度

package main

func main() {
    a := [2]int{}
    b := [...][2]int{
        {10, 20},
        {30, 40},
        {50, 60},
    }

    println(len(a), cap(a))   // 2 2
    println(len(b), cap(b))   // 3 3
    println(len(b[1]), cap(b[1]))  // 2 2
}

2.5 使用 range 迭代数组
for 循环可用于遍历数组中的元素。

package main

import "fmt"

func main() {
    a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(a); i++ { // looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}

上面的程序使用 for 循环遍历数组中的元素，从索引 0 到 length of the array - 1

Go 提供了一种更好、更简洁的方法，通过使用 for 循环的 range 方法来遍历数组。range返回索引和该索引处的值。让我们使用 range 重写上面的代码。我们还可以获取数组中所有元素的总和。

package main

import "fmt"

func main() {
    a := [...]float64{67.7, 89.8, 21, 78}
    sum := float64(0)
    for i, v := range a { //range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a", sum)
}

上述程序的第 8 行 for i, v := range a 利用的是 for 循环 range 方式。 它将返回索引和该索引处的值。 我们打印这些值，并计算数组 a 中所有元素的总和。

如果你只需要值并希望忽略索引，则可以通过用 _ 空白标识符替换索引来执行。

for _, v := range a { // ignores index  }
1
上面的 for 循环忽略索引，同样值也可以被忽略。

2.6 数组操作符操作
如元素类型支持”==,!=”操作符，那么数组也支持此操作

package main

func main() {
    var a, b [2]int
    println(a == b)  //true

    c := [2]int{1, 2}
    d := [2]int{0, 1}
    println(c != d) //true
    /*
        var e, f [2]map[string]int
        println(e == f)  //invalid operation: e == f ([2]map[string]int cannot be compared)
    */
}

3 数组高级用法
3.1 多维数组
到目前为止我们创建的数组都是一维的，Go 语言可以创建多维数组。

package main

import (
    "fmt"
)

func printArray(a [3][2]string) {
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func main() {
    a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
    }
    printArray(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printArray(b)
}

/*output
lion tiger 
cat dog 
pigeon peacock 

apple samsung 
microsoft google 
AT&T T-Mobile 
*/

多维数组遍历

package main

import (
    "fmt"
)

func main() {
    var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
    for k1, v1 := range f {
        for k2, v2 := range v1 {
            fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
        }
        fmt.Println()
    }
}
/*
output:
(0,0)=1 (0,1)=2 (0,2)=3 
(1,0)=7 (1,1)=8 (1,2)=9 
 */

3.2 数组指针和指针数组
要分清指针数组和数组指针的区别。指针数组是指元素为指针类型的数组，数组指针是获取数组变量的地址。

package main

import "fmt"

func main() {
    x, y := 10, 20
    a := [...]*int{&x, &y}
    p := &a

    fmt.Printf("%T,%v\n", a, a)  //[2]*int,[0xc042062080 0xc042062088]
    fmt.Printf("%T,%v\n", p, p)  //*[2]*int,&[0xc042062080 0xc042062088]
}

可获取任意元素地址

func main() {
    a := [...]int{1, 2}
    println(&a, &a[0], &a[1])  //0xc042049f68 0xc042049f68 0xc042049f70
}

数组指针可以直接用来操作元素

func main() {
    a := [...]int{1, 2}
    p := &a

    p[1] += 10
    println(p[1])   //12
}

4 数组使用常见坑
定义数组类型时，数组长度必须是非负整型常量表达式，长度是类型组成部分。也就是说，元素类型相同，但长度不同的数组不属于同一类型。

例子：

func main() {
    var d1 [3]int
    var d2 [2]int
    d1 = d2 //cannot use d2 (type [2]int) as type [3]int in assignment
}

5 数组总结
数组：是同一种数据类型的固定长度的序列。

数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。

长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。

数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1。数组索引常用操作如下：

for i := 0; i < len(a); i++ {
   ...
} 

for index, v := range a {
   ...
} 

访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic

数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。

支持 “==”、”!=” 操作符，因为内存总是被初始化过的。

指针数组 [n]*T，数组指针*[n]T。



----------------------------------------------------
golang中的三个点 '...' 的用法
----------------------------------------------------


'...' 其实是go的一种语法糖。 
它的第一个用法主要是用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。 
第二个用法是slice可以被打散进行传递。

下面直接上例子：

func test1(args ...string) { //可以接受任意个string参数
    for _, v:= range args{
        fmt.Println(v)
    }
}

func main(){
var strss= []string{
        "qwr",
        "234",
        "yui",
        "cvbc",
    }
    test1(strss...) //切片被打散传入
}

结果：
qwr
234
yui
cvbc
其中strss切片内部的元素数量可以是任意个，test1函数都能够接受
--------------------- 
第二个例子：
  var strss= []string{
        "qwr",
        "234",
        "yui",

    }
    var strss2= []string{
        "qqq",
        "aaa",
        "zzz",
        "zzz",
    }
strss=append(strss,strss2...) //strss2的元素被打散一个个append进strss
fmt.Println(strss)
结果：

[qwr 234 yui qqq aaa zzz zzz]
1
如果没有’…’，面对上面的情况，无疑会增加代码量，有了’…’，是不是感觉简洁了许多
--------------------- 

