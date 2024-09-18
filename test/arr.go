package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 首字母大写的标识符是可导出的（public），可以从其他包中访问。
// 首字母小写的标识符是不可导出的（private），只能在当前包内部访问。
func testArr() {
	//指定数组大小
	arr := [10]string{"a", "b", "c", "d", "e"}
	//自动确定数组大小
	arr2 := [...]string{"a", "b", "c", "d", "e"}

	//通过下标截取
	slice3 := arr[0:3]
	//字面量初始化切片
	slice := []string{"a", "b", "c", "d"}
	//make关键字创建切片
	slice2 := make([]string, 10)

	fmt.Println(arr, arr2, slice, slice2, slice3)
}
func Test() {
	//指定数组大小
	arr := [10]string{"a", "b", "c", "d", "e"}
	//自动确定数组大小
	arr2 := [...]string{"a", "b", "c", "d", "e"}

	//通过下标截取
	slice3 := arr[0:3]
	//字面量初始化切片
	slice := []string{"a", "b", "c", "d"}
	//make关键字创建切片
	slice2 := make([]string, 10)

	fmt.Println(arr, arr2, slice, slice2, slice3)
}
func Lianjie() {
	// 创建数据库连接
	db, err := sql.Open("mysql", "root:131415@tcp(localhost:3306)/myblog")
	if err != nil {
		// 处理连接错误
	}
	defer db.Close()
	// 执行查询
	rows, err := db.Query("SELECT username FROM tb_user")
	if err != nil {
		// 处理查询错误
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			// 处理扫描错误
		}
		// 处理查询结果
		fmt.Printf("Name: %s", name)
	}
}
