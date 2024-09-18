package main

import "fmt"

type student struct {
	age  int
	name string
}

// 匿名结构体
var dog struct {
	id   int
	name string
}

func (stu student) study() {
	fmt.Println(stu.name + "学习！")
}
func (student student) getName() string {
	return student.name
}
func (student *student) setName(name string) {
	student.name = name
}

func main() {
	//类型定义
	type Myint int
	type Mystring = string

	dog.id = 1
	dog.name = "歪歪"
	fmt.Println(dog)
	fmt.Println(dog.name)

	var zhangsan = new(student)
	zhangsan.age = 10
	fmt.Println(*zhangsan)
	zhangsan.study()
	zhangsan.setName("张三")
	fmt.Println(zhangsan.getName())
}
