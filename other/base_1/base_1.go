package main

import (
	"fmt"
)

func main() {
	var a [4]int          //元素自动初始化为零[0 0 0 0]
	b := [4]int{2, 5}     //未提供初始化值得元素自动初始化为0  [2 5 0 0]
	c := [4]int{5, 3: 10} //可指定索引位置初始化 [5 0 0 10]; 数组下标为3的值设置为10

	fmt.Println(typeof(c))

	d := [...]int{1, 2, 3}    //编译器按初始化值数量确定数组长度 [1 2 3]
	e := [...]int{10, 4: 100} //支持索引初始化，但注意数组长度与此有关 [10 0 0 100]
	fmt.Println(a, b, c, d, e)
}

func typeof(c interface{}) string {
	return fmt.Sprintf("%T", c)
}
