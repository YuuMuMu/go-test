package main

import "fmt"

var i = varInit()

func varInit() int {
	fmt.Println("var init")
	return 1
}
func init() {
	fmt.Println("init")
}

//func main() {
//	fmt.Println("main")
//
//	//var init
//	//init
//	//main
//
//}
