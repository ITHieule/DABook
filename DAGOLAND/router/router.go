package router

import (
	"DAGOLAND/handler"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Đảm bảo rằng các route này đã được đăng ký
	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/todos", handler.Gettodo).Methods("GET")            // Lấy danh sách
	r.HandleFunc("/todos", handler.AddTodo).Methods("POST")           // Thêm công việc
	r.HandleFunc("/todos/{id}", handler.UpdateTodo).Methods("PUT")    // Sửa công việc
	r.HandleFunc("/todos/{id}", handler.DeleteTodo).Methods("DELETE") // Xóa công việc
	return r
}
