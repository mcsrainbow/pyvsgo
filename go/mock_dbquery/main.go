package main

import (
	"fmt"
)

// Database 接口定义
type Database interface {
	Query(query string) (result string, err error)
}

// MockDatabase 结构体实现 Database 接口
type MockDatabase struct{}

func (m MockDatabase) Query(query string) (result string, err error) {
	return "mock result", nil
}

// RealDatabase 结构体实现 Database 接口
type RealDatabase struct{}

func (r RealDatabase) Query(query string) (result string, err error) {
	return "real result", nil
}

// 使用 Database 接口的函数
func ExecuteQuery(db Database, query string) {
	result, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func main() {
	var db Database

	// 使用 MockDatabase
	db = MockDatabase{}
	ExecuteQuery(db, "SELECT * FROM table") // Output: mock result

	// 使用 RealDatabase
	db = RealDatabase{}
	ExecuteQuery(db, "SELECT * FROM table") // Output: real result
}
