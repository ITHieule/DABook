package tasks

import (
	"fmt"
	"github.com/robfig/cron"
	"web-api/internal/api/services"
)

func StartScheduler() {
	// Tạo đối tượng cron
	c := cron.New()

	// Lấy lịch trình từ file cấu hình
	schedule := "0 */5 * * * *" // Chạy mỗi 5 phút, bắt đầu tại giây thứ 0

	// Thêm task với lịch trình
	err := c.AddFunc(schedule, func() {
		// Gọi CheckBookCountAndNotify và xử lý lỗi
		if err := services.CheckBookCountAndNotify(); err != nil {
			fmt.Println("Error in CheckBookCountAndNotify:", err)
		}
	})
	if err != nil {
		fmt.Println("Error scheduling task:", err)
		return
	}

	// Bắt đầu scheduler
	c.Start()

	fmt.Println("Scheduler started. Running every phut.")

	// Giữ chương trình hoạt động
	select {}
}
