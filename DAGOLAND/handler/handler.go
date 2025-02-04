package handler

import (
	"DAGOLAND/database"
	"DAGOLAND/model"
	"DAGOLAND/until"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Decode JSON từ body
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Dữ liệu không hợp lệ", http.StatusBadRequest)
		return
	}
	// Mã hóa mật khẩu
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Lỗi mã hóa mật khẩu", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword
	// Kiểm tra dữ liệu đầu vào
	if user.Username == "" || user.Password == "" || user.Email == "" {
		http.Error(w, "Dữ liệu không đầy đủ", http.StatusBadRequest)
		return
	}

	// Kết nối cơ sở dữ liệu
	database.ConnecDB() // Kết nối CSDL (hàm này không nhận tham số)

	// Thêm người dùng vào cơ sở dữ liệu
	query := "INSERT INTO user (username, email, password) VALUES (?, ?, ?)"
	result, err := database.DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Không thể đăng ký người dùng", http.StatusInternalServerError)
		return
	}

	// Lấy ID người dùng vừa tạo
	lastInsertID, _ := result.LastInsertId()
	user.ID = uint(lastInsertID)

	// Trả về phản hồi
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, req *http.Request) {
	var user model.User
	var input model.User
	json.NewDecoder(req.Body).Decode(&input)
	// Truy vấn để lấy thông tin người dùng từ cơ sở dữ liệu (không cần mật khẩu)
	query := "SELECT id, username, email, password FROM user WHERE username = ?"
	err := database.DB.QueryRow(query, input.Username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Failed to login", http.StatusInternalServerError)
		return
	}

	// So sánh mật khẩu đã mã hóa với mật khẩu người dùng nhập vào
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Tạo JWT token
	token, err := until.GenerateJWT(user.ID) // Chú ý: Bạn có thể thêm quyền truy cập (role) vào JWT nếu cần
	if err != nil {
		http.Error(w, "Lỗi khi tạo token", http.StatusInternalServerError)
		return
	}

	// Gửi token cho người dùng
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})

}
func Gettodo(w http.ResponseWriter, r *http.Request) {

	rows, err := database.DB.Query("SELECT id, user_id, title, description, completed, created_at, updated_at,image_url FROM todo ")
	if err != nil {
		http.Error(w, "Failed to retrieve todos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []model.TodoList
	for rows.Next() {
		var todo model.TodoList
		err = rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.Image)
		if err != nil {
			http.Error(w, "Failed to parse todos", http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

//func GetTodoImage(w http.ResponseWriter, r *http.Request) {
//	// Lấy ID công việc từ URL
//	id := mux.Vars(r)["id"]
//
//	// Lấy đường dẫn đến hình ảnh từ cơ sở dữ liệu
//	var imagePath string
//	err := database.DB.QueryRow("SELECT image_url FROM todo WHERE id = ?", id).Scan(&imagePath)
//	if err != nil || imagePath == "" {
//		http.Error(w, "Image not found", http.StatusNotFound)
//		return
//	}
//
//	// Đọc hình ảnh từ đường dẫn
//	file, err := os.Open(imagePath)
//	if err != nil {
//		http.Error(w, "Failed to open image", http.StatusInternalServerError)
//		return
//	}
//	defer file.Close()
//
//	// Lấy thông tin tệp hình ảnh
//	fileInfo, err := file.Stat()
//	if err != nil {
//		http.Error(w, "Failed to get file info", http.StatusInternalServerError)
//		return
//	}
//
//	// Xác định loại hình ảnh (ví dụ: JPEG, PNG)
//	contentType := "image/jpeg" // Giả sử hình ảnh là JPEG
//	if ext := filepath.Ext(imagePath); ext == ".png" {
//		contentType = "image/png"
//	}
//
//	// Thiết lập tiêu đề phản hồi
//	w.Header().Set("Content-Type", contentType)
//	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
//
//	// Gửi tệp hình ảnh tới client
//	http.ServeFile(w, r, imagePath)
//}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.TodoList
	json.NewDecoder(r.Body).Decode(&todo)

	if todo.UserID == 0 || todo.Title == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO todo (user_id, title, description, completed, image_url) VALUES (?, ?, ?, ?, ?)"
	result, err := database.DB.Exec(query, todo.UserID, todo.Title, todo.Description, todo.Completed, todo.Image)
	if err != nil {
		http.Error(w, "Failed to add todo", http.StatusInternalServerError)
		return
	}

	lastInsertID, _ := result.LastInsertId()
	todo.ID = uint(lastInsertID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var todo model.TodoList
	json.NewDecoder(r.Body).Decode(&todo)

	query := "UPDATE todo SET title = ?, description = ?, completed = ?, image_url = ? WHERE id = ?"
	_, err := database.DB.Exec(query, todo.Title, todo.Description, todo.Completed, todo.Image, id)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo updated"})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := "DELETE FROM todo WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted"})
}

// Hàm mã hóa mật khẩu
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

//type key string
//
//const userKey key = "user"
//
//// Middleware để thêm thông tin user vào context
//func AdminOnlyHandler(w http.ResponseWriter, r *http.Request) {
//	// Sử dụng key chính xác để lấy thông tin từ context
//	claims, ok := r.Context().Value(userKey).(map[string]interface{})
//	if !ok || claims == nil {
//		http.Error(w, "Unauthorized", http.StatusUnauthorized)
//		return
//	}
//
//	// Kiểm tra vai trò
//	role, ok := claims["role"].(string)
//	if !ok || role != "admin" {
//		http.Error(w, "Forbidden", http.StatusForbidden)
//		return
//	}
//
//	// Nếu là admin
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte("Welcome, Admin!"))
//}
