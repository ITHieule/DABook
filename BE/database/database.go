package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// Hàm kết nối tới cơ sở dữ liệu
func ConnecDB() {
	// Chuỗi kết nối (Data Source Name - DSN)
	server := "192.168.40.37"
	port := 1433
	user := "sa"
	password := "jack"
	database := "TB_ERP"

	// Tạo chuỗi kết nối DSN
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable&trustServerCertificate=true",
		user, password, server, port, database)

	// Mở kết nối tới cơ sở dữ liệu
	var err error
	DB, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal("Failed to create database connection:", err)
	}

	// Kiểm tra kết nối tới cơ sở dữ liệu
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	rows, err := DB.Query("SELECT XieXing, SheHao FROM xxzl")
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	defer rows.Close() // Đảm bảo đóng kết nối sau khi đọc dữ liệu

	// Giả sử bảng xxzl có các cột: ID (int), XieXing (string), SheHao (string), và một cột khác là Name (string)
	for rows.Next() {
		var XieXing, SheHao string

		// Quét dữ liệu từ hàng hiện tại vào các biến
		err := rows.Scan(&XieXing, &SheHao)
		if err != nil {
			log.Println("Failed to scan row:", err)
			continue
		}

		// In dữ liệu đã quét ra màn hình
		log.Printf("XieXing: %s, SheHao: %s", XieXing, SheHao)
	}

	// Kiểm tra lỗi sau khi lặp qua rows
	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}

	log.Println("Connected to SQL Server successfully!")
}
