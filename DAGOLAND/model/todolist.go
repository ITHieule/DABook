package model

type TodoList struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"` // Khóa ngoại
	Title       string `gorm:"not null"`
	Description string
	Completed   bool   `gorm:"default:false"`
	Image       string `gorm:"not null"`
	CreatedAt   string `gorm:"autoCreateTime"`
	UpdatedAt   string `gorm:"autoUpdateTime"`
}
