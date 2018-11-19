如果我们要遍历某个数组，Map集合，Slice切片等，Go语言(Golang)为我们提供了比较好用的For Range方式。range是一个关键字，表示范围，和for配合使用可以迭代数组,Map等集合。它的用法简洁，而且map、channel等也都是用for range的方式，所以在编码中我们使用for range进行循环迭代是最多的。对于这种最常使用的迭代，尤其是和for i=0;i<N;i++对比，性能怎么样？我们进行下示例分析，让我们对for range循环有个更深的理解，便于我们写出性能更高的程序。

基本用法
for range的使用非常简单，这里演示下两种集合类型的使用。

package main

import "fmt"

func main() {
    ages:=[]string{"10", "20", "30"}

    for i,age:=range ages{
        fmt.Println(i,age)
    }
}
这是针对 Slice 切片的迭代使用,使用range关键字返回两个变量i,age，第一个是 Slice 切片的索引，第二个是 Slice 切片中的内容，所以我们打印出来：

0 10
1 20
2 30
关于Go语言 Slice 切片的，可以参考我以前写的这篇 Go语言实战笔记（五）| Go 切片

下面再看看map（字典）的for range使用示例。

package main

import "fmt"

func main() {
    ages:=map[string]int{"张三":15,"李四":20,"王武":36}

    for name,age:=range ages{
        fmt.Println(name,age)
    }
}
在使用for range迭代map的时候，返回的第一个变量是key,第二个变量是value，也就是我们例子中对应的name和ages。我们运行程序看看输出结果。

张三 15
李四 20
王武 36
这里需要注意的是，for range map返回的K-V键值对顺序是不固定的，是随机的，这次可能是张三-15第一个出现，下一次运行可能是王武-36第一个被打印了。
关于Map更详细的可以参考我以前的一篇文章 Go语言实战笔记（六）| Go Map。

常规for循环对比
比如对于 Slice 切片，我们有两种迭代方式：一种是常规的for i:=0;i<N;i++的方式；一种是for range的方式，下面我们看看两种迭代的性能。

func ForSlice(s []string) {
    len := len(s)
    for i := 0; i < len; i++ {
        _, _ = i, s[i]
    }
}

func RangeForSlice(s []string) {
    for i, v := range s {
        _, _ = i, v
    }
}
为了测试，写了这两种循环迭代 Slice 切片的函数，从实现上看，他们的逻辑是一样的，保证我们可以在同样的情况下测试。

import "testing"

const N  =  1000

func initSlice() []string{
    s:=make([]string,N)
    for i:=0;i<N;i++{
        s[i]="www.flysnow.org"
    }
    return s;
}

func BenchmarkForSlice(b *testing.B) {
    s:=initSlice()

    b.ResetTimer()
    for i:=0; i<b.N;i++  {
        ForSlice(s)
    }
}

func BenchmarkRangeForSlice(b *testing.B) {
    s:=initSlice()

    b.ResetTimer()
    for i:=0; i<b.N;i++  {
        RangeForSlice(s)
    }
}
这事Bench基准测试的用例，都是在相同的情况下，模拟长度为1000的 Slice 切片的遍历。然后我们运行go test -bench=. -run=NONE查看性能测试结果。

BenchmarkForSlice-4              5000000    287 ns/op
BenchmarkRangeForSlice-4         3000000    509 ns/op
从性能测试可以看到，常规的for循环，要比for range的性能高出近一倍，到这里相信大家已经知道了原因，没错，因为for range每次是对循环元素的拷贝，所以集合内的预算越复杂，性能越差，而反观常规的for循环，它获取集合内元素是通过s[i]，这种索引指针引用的方式，要比拷贝性能要高的多。

既然是元素拷贝的问题，我们迭代 Slice 切片的目的也是为了获取元素，那么我们换一种方式实现for range。

func RangeForSlice(s []string) {
    for i, _ := range s {
        _, _ = i, s[i]
    }
}
现在，我们再次进行 Benchmark 性能测试,看看效果。

BenchmarkForSlice-4              5000000    280 ns/op
BenchmarkRangeForSlice-4         5000000    277 ns/op
恩，和我们想的一样，性能上来了，和常规的for循环持平了。原因就是我们通过_舍弃了元素的复制，然后通过s[i]获取迭代的元素，既提高了性能，又达到了目的。

Map 遍历
对于Map来说，我们并不能使用for i:=0;i<N;i++的方式，当然如果你有全部的key元素列表除外，所以大部分情况下我们都是使用for range的方式。

func RangeForMap1(m map[int]string) {
    for k, v := range m {
        _, _ = k, v
    }
}

const N = 1000

func initMap() map[int]string {
    m := make(map[int]string, N)
    for i := 0; i < N; i++ {
        m[i] = fmt.Sprint("www.flysnow.org",i)
    }
    return m
}

func BenchmarkRangeForMap1(b *testing.B) {
    m:=initMap()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        RangeForMap1(m)
    }
}
http://www.flysnow.org/

飞雪无情的博客

以上示例是map遍历的函数以及benchmark测试，我都写在一起了，运行测试看一下效果。

BenchmarkForSlice-8              5000000    298 ns/op
BenchmarkRangeForSlice-8         3000000    475 ns/op
BenchmarkRangeForMap1-8           100000    14531 ns/op
相比 Slice 来说，Map的遍历的性能更差，可以说是惨不忍睹。好，我们开始下优化，思路也是减少值得拷贝。测试中的RangeForSlice也慢的原因是我把RangeForSlice还原成了值得拷贝，以便于对比性能。

func RangeForMap2(m map[int]string) {
    for k, _ := range m {
        _, _ = k, m[k]
    }
}

func BenchmarkRangeForMap2(b *testing.B) {
    m := initMap()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        RangeForMap2(m)
    }
}
再次运行下性能测试看下效果。

BenchmarkForSlice-8              5000000       298 ns/op
BenchmarkRangeForSlice-8         3000000       475 ns/op
BenchmarkRangeForMap1-8           100000       14531 ns/op
BenchmarkRangeForMap2-8           100000       23199 ns/op
额，是不是发现点不对，方法BenchmarkRangeForMap2的性能明显下降了，这个可以从每次操作的耗时看出来(虽然性能测试秒执行的次数还是一样)。和我们上面测试的Slice不一样，这次不止没有提升，反而下降了。

继续修改Map2函数的实现为：

func RangeForMap2(m map[int]Person) {
    for  range m {
    }
}
什么都不做，只迭代，再次运行性能测试。

BenchmarkForSlice-8              5000000       301 ns/op
BenchmarkRangeForSlice-8         3000000       478 ns/op
BenchmarkRangeForMap1-8           100000       14822 ns/op
BenchmarkRangeForMap2-8           100000       14215 ns/op
*我们惊奇的发现，什么都不做，和获取K-V值的操作性能是一样的，和Slice完全不一样，不是说 for range值拷贝损耗性能呢？都哪去了？大家猜一猜，可以结合下一节的原理实现

for range 原理
通过查看https://github.com/golang/gofrontend源代码，我们可以发现for range的实现是：

// Arrange to do a loop appropriate for the type.  We will produce
  //   for INIT ; COND ; POST {
  //           ITER_INIT
  //           INDEX = INDEX_TEMP
  //           VALUE = VALUE_TEMP // If there is a value
  //           original statements
  //   }
并且对于Slice,Map等各有具体不同的编译实现,我们先看看for range slice的具体实现

  // The loop we generate:
  //   for_temp := range
  //   len_temp := len(for_temp)
  //   for index_temp = 0; index_temp < len_temp; index_temp++ {
  //           value_temp = for_temp[index_temp]
  //           index = index_temp
  //           value = value_temp
  //           original body
  //   }
先是对要遍历的 Slice 做一个拷贝，获取长度大小，然后使用常规for循环进行遍历，并且返回值的拷贝。

再看看for range map的具体实现：

  // The loop we generate:
  //   var hiter map_iteration_struct
  //   for mapiterinit(type, range, &hiter); hiter.key != nil; mapiternext(&hiter) {
  //           index_temp = *hiter.key
  //           value_temp = *hiter.val
  //           index = index_temp
  //           value = value_temp
  //           original body
  //   }
也是先对map进行了初始化，因为map是*hashmap，所以这里其实是一个*hashmap指针的拷贝。

结合着这两个具体的for range编译器实现，可以看看为什么for range slice的_优化方式有用，而for range map的方式没用呢？欢迎大家留言回答。

本文为原创文章，转载自: 飞雪无情,
以下是该作者的官网:
http://www.flysnow.org/