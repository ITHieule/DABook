package until

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Khóa bí mật để ký token
var jwtSecret = []byte("your_secret_key")

// Hàm tạo JWT token
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(), // Hạn token 2 giờ
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
