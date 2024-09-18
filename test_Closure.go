package main

import "fmt"

func test() {
	// 外部函数
	outerFunc := func() func(int) int {
		// 外部函数中定义了一个局部变量
		sum := 0

		// 返回一个闭包  闭包同外部函数类型一样
		return func(x int) int {
			sum += x
			return sum
		}
	}

	// 创建一个闭包
	myClosure := outerFunc()

	// 使用闭包
	fmt.Println(myClosure(1)) // 输出: 1
	fmt.Println(myClosure(2)) // 输出: 3
	fmt.Println(myClosure(3)) // 输出: 6
}

//func main() {
//	test()
//}
