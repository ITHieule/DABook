package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/types"
)

type BookService struct {
	*BaseService
}

var Book = &BookService{}

func (s *BookService) GetDataService() ([]types.Booktypes, error) {
	var Book []types.Booktypes
	// kết nối cơ sở dữ liệu
	db, err := database.DB1Connection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()
	query := `
		SELECT  ID,Ten,IDloai,Soluong
		FROM Book		
	`

	err = db.Raw(query).Scan(&Book).Error

	if err != nil {
		fmt.Println("Query execution error:", err)
	}

	return Book, nil
}
func (s *BookService) AddDataService(requestParams *request.Book) (types.Booktypes, error) {
	var Book types.Booktypes
	// kết nối cơ sở dữ liệu
	db, err := database.DB1Connection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return Book, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()
	query := `
        INSERT INTO Book (
           Ten,IDloai,Soluong
        ) VALUES (?, ?, ?)
    `
	// Truyền tham số vào câu truy vấn
	err = db.Raw(query,
		requestParams.Ten,
		requestParams.IDloai,
		requestParams.Soluong,
	).Scan(&Book).Error

	if err != nil {
		fmt.Println("Query execution error:", err)
	}
	return Book, nil
}
func (s *BookService) DeleteDataService(ID int) {

	// kết nối cơ sở dữ liệu
	db, err := database.DB1Connection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()
	query := `
       DELETE FROM Book
		WHERE ID = ?
    `
	// Truyền tham số vào câu truy vấn
	err = db.Raw(query, ID).Scan(&ID).Error
	if err != nil {
		fmt.Println("Query execution error:", err)
	}
	return
}
func (s *BookService) UpdateDataService(requestParams *request.Book) (types.Booktypes, error) {
	var Book types.Booktypes
	// kết nối cơ sở dữ liệu
	db, err := database.DB1Connection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return Book, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()
	query := `
    UPDATE Book
    SET Ten = ?, IDloai = ?, Soluong = ?
    WHERE ID = ?
`
	err = db.Exec(query,
		requestParams.Ten,
		requestParams.IDloai,
		requestParams.Soluong,
		requestParams.ID,
	).Error

	if err != nil {
		fmt.Println("Query execution error:", err)
		return Book, err
	}

	return Book, nil
}
func (s *BookService) OderDataService(requestParams *request.Oder) (types.Odertypes, error) {
	var oder types.Odertypes

	// Kết nối cơ sở dữ liệu
	db, err := database.DB1Connection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return oder, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()

	// Bắt đầu một transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Println("Transaction rollback due to panic")
		}
	}()

	// Câu truy vấn cập nhật số lượng sách
	updateBookQuery := `
	UPDATE Book
	SET Soluong = Soluong - ?
	WHERE ID = ?
	AND Soluong >= ?;  -- Đảm bảo số lượng còn đủ để giảm
	`

	err = tx.Exec(updateBookQuery,
		requestParams.Soluong, // Số lượng giảm
		requestParams.IDbook,  // ID sách
		requestParams.Soluong, // Kiểm tra không giảm quá số lượng
	).Error
	if err != nil {
		fmt.Println("Failed to update book quantity:", err)
		tx.Rollback()
		return oder, err
	}

	// Câu truy vấn thêm đơn hàng mới
	insertOrderQuery := `
	INSERT INTO Oder
		(Ngaydathang, IDbook, IDkhachhang, Soluong)
	VALUES
		(?, ?, ?, ?);
	`

	err = tx.Exec(insertOrderQuery,
		requestParams.Ngaydathang, // Ngày đặt hàng
		requestParams.IDbook,      // ID sách
		requestParams.IDkhachhang, // ID khách hàng
		requestParams.Soluong,     // Số lượng
	).Error
	if err != nil {
		fmt.Println("Failed to insert order:", err)
		tx.Rollback()
		return oder, err
	}

	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		fmt.Println("Transaction commit error:", err)
		return oder, err
	}

	return oder, nil
}
func (s *BookService) SearchIDService(requestParams request.Book) ([]types.Booktypes, error) {
	var Book []types.Booktypes

	// Kết nối cơ sở dữ liệu
	db, err := database.DB1Connection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn với tham số
	query := `
		SELECT *
		FROM Book
		WHERE IDloai = ? OR Ten = ? OR ID = ?
	`

	// Thực hiện truy vấn và ánh xạ kết quả vào cấu trúc dữ liệu
	err = db.Raw(query,
		requestParams.IDloai,
		requestParams.Ten,
		requestParams.ID,
	).Scan(&Book).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}

	return Book, nil
}

// Hàm mã hóa mật khẩu
func HashPassword(Password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
