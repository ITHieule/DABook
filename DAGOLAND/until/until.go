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

//func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		// Kiểm tra loại signing method
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, jwt.ErrSignatureInvalid
//		}
//		return jwtSecret, nil
//	})
//
//	if err != nil || !token.Valid {
//		return nil, err
//	}
//
//	// Trả về claims nếu token hợp lệ
//	if claims, ok := token.Claims.(jwt.MapClaims); ok {
//		return claims, nil
//	}
//	return nil, jwt.ErrInvalidKey
//}
