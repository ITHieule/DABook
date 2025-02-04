package services

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"web-api/internal/pkg/database"
)

// Hàm gửi email
func SendEmail(to string, subject string, body string) error {

	viper.SetConfigName("config")   // Tên file (không bao gồm đuôi .yaml)
	viper.SetConfigType("yaml")     // Loại file
	viper.AddConfigPath("././data") // Đường dẫn file (thư mục hiện tại)

	// Đọc file cấu hình
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Lấy các thông tin cấu hình từ biến môi trường
	emailHost := viper.GetString("email.host")
	emailPort := viper.GetString("email.port")
	emailUsername := viper.GetString("email.username")
	emailPassword := viper.GetString("email.password")

	// Cấu hình email
	mail := gomail.NewMessage()
	mail.SetHeader("From", emailUsername)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", body)

	// Cấu hình SMTP server
	portNumber, _ := strconv.Atoi(emailPort) // Chuyển port sang kiểu số
	dialer := gomail.NewDialer(emailHost, portNumber, emailUsername, emailPassword)

	// Gửi email
	err := dialer.DialAndSend(mail)
	if err != nil {
		return err
	}
	return nil
}

func CheckBookCountAndNotify() error {

	// Kết nối cơ sở dữ liệu
	db, err := database.DB1Connection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()

	// Kiểm tra xem có loại sách nào có số lượng bằng 0
	var bookTitles []string
	query := "SELECT Ten FROM Book WHERE Soluong = 0"
	err = db.Raw(query).Scan(&bookTitles).Error
	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}
	viper.SetConfigName("config")   // Tên file (không bao gồm đuôi .yaml)
	viper.SetConfigType("yaml")     // Loại file
	viper.AddConfigPath("././data") // Đường dẫn file (thư mục hiện tại)

	// Đọc file cấu hình
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	// Nếu có loại sách nào đó có số lượng bằng 0, gửi email
	if len(bookTitles) > 0 {
		// Lấy email nhận thông báo từ file cấu hình
		emailUsername := viper.GetString("email.username")

		subject := "Thông báo: Một số loại sách trong hệ thống có số lượng bằng 0"
		body := "Hệ thống thông báo rằng các loại sách sau đây hiện có số lượng bằng 0:\n\n"

		// Thêm danh sách sách vào nội dung email
		for _, title := range bookTitles {
			body += "- " + title + "\n"
		}

		body += "\nVui lòng kiểm tra hệ thống và cập nhật thêm sách."

		// Gửi email
		err = SendEmail(emailUsername, subject, body)
		if err != nil {
			fmt.Println("Error sending email:", err)
			return err
		}

		fmt.Println("Email sent successfully to", emailUsername)
	} else {
		fmt.Println("Tất cả các loại sách đều có số lượng > 0.")
	}

	return nil
}
