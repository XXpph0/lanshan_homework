package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 定义一个全局对象db
var db *sqlx.DB

func initDB() {
	var err error

	dsn := "username:password@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("连接数据库失败，错误：%v\n", err)
		return
	}

	fmt.Println("连接到MySQL数据库...")
	return
}

func insertData() {
	// 向student表中插入十条记录
	for i := 1; i <= 10; i++ {
		result, err := db.Exec("INSERT INTO student (name) VALUES (?)", fmt.Sprintf("Student%d", i))
		if err != nil {
			fmt.Printf("插入数据失败，错误：%v\n", err)
			continue
		}
		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("成功插入%d条数据\n", rowsAffected)
	}
}

func readData() {
	// 读取并打印student表中的所有记录
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		fmt.Printf("读取数据失败，错误：%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Printf("扫描数据失败，错误：%v\n", err)
			continue
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}

func main() {
	initDB()
	insertData()
	readData()
}
