package middleware

//
//import (
//	"DAGOLAND/until"
//	"context"
//	"net/http"
//	"strings"
//)
//
//type key string
//
//const userKey key = "user"
//
//func AuthMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// Lấy token từ header Authorization
//		authHeader := r.Header.Get("Authorization")
//		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
//			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
//			return
//		}
//
//		// Bỏ tiền tố "Bearer "
//		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
//
//		// Validate token
//		claims, err := until.ValidateJWT(tokenString)
//		if err != nil {
//			http.Error(w, "Invalid token", http.StatusUnauthorized)
//			return
//		}
//
//		// Gắn thông tin user vào context
//		ctx := context.WithValue(r.Context(), userKey, claims)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
