package main

import (
	"fmt"
	"math"
)

// 接受一个可变数量的空接口类型参数（interface{}）。空接口类型可以容纳任何类型的值
func myfunc(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
func test2() {
	//匿名函数由一个不带函数名的函数声明和函数体组成。
	//匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

	//自己调用
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}(8)

	//单独调用
	//fmt.Println(getSqrt(8))
	fmt.Println(getSqrt)
	myfunc(1, 1, "1231")
}

//func main() {
//
//}

func chushihua() {
	//初始化方式一
	hash := map[string]int{
		"1":    2,
		"lisi": 2,
	}
	//初始化方式二
	mp := make(map[string]string)
	mp["lisi"] = "1231"

	fmt.Println(hash)
	for key, value := range hash {
		fmt.Println(key, " ----> ", value)
	}
	//删除
	delete(hash, "1")
	fmt.Println(hash)
}

func add(a int, b int) int {
	return a + b
}
