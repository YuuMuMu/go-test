package main

import "fmt"

func test1() {
	var whatever [5]struct{}
	for i := range whatever {
		//由于闭包用到的变量 i 在执行的时候已经变成4
		defer func() { fmt.Println(i) }() // 4 4 4 4 4
	}

}

//func main() {
//	test1()
//}
