package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"web-api/internal/api/until"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/types"
)

type UserService struct {
	*BaseService
}

var User = &UserService{}

func (s *UserService) Register(requestParams *request.User) ([]types.Usertypes, error) {
	var User []types.Usertypes
	// kết nối cơ sở dữ liệu
	db, err := database.DB1Connection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	// Mã hóa mật khẩu
	hashedPassword, err := HashPassword(requestParams.Password)
	if err != nil {

		return nil, err
	}
	requestParams.Password = hashedPassword
	// Kiểm tra dữ liệu đầu vào
	dbInstance, _ := db.DB()
	defer dbInstance.Close()
	query := "INSERT INTO User (Name, Email, Password) VALUES (?, ?, ?)"
	err = db.Raw(query,
		requestParams.Name,
		requestParams.Email,
		requestParams.Password,
	).Scan(&User).Error
	if err != nil {
		fmt.Println("Query execution error:", err)
	}

	return User, nil
}
func (s *UserService) Login(requestParams *request.User) (string, error) {
	var user types.Usertypes

	// Kết nối cơ sở dữ liệu
	db, err := database.DB1Connection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return "", err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()

	// Truy vấn thông tin người dùng dựa trên email
	query := "SELECT * FROM User WHERE Email = ?"
	err = db.Raw(query, requestParams.Email).Scan(&user).Error

	if err != nil {
		fmt.Println("Query execution error:", err)
		return "", err
	}

	// So sánh mật khẩu đã mã hóa với mật khẩu người dùng nhập vào
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestParams.Password)); err != nil {
		return "", nil
	}

	// Tạo JWT token
	token, err := until.GenerateJWT(user.ID)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return token, nil
	}

	// Trả về thông tin người dùng và token
	return token, nil

}
