package main

import (
	"DAGOLAND/database"
	"DAGOLAND/router"
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//var todo = []string{"Đi Học"}
	//var check int
	//for {
	//	fmt.Println("1.Xem danh sách công việc")
	//	fmt.Println("2.Thêm công việc")
	//	fmt.Println("3.Thoát")
	//	fmt.Println("Nhập lựa chọn")
	//	fmt.Scan(&check)
	//	if check < 1 || check > 5 {
	//		fmt.Println("Lựa chọn không hợp lệ. Vuii lòng nhập lại: ")
	//		continue
	//	}
	//	switch check {
	//	case 1:
	//		if len(todo) == 0 {
	//			fmt.Println("không có công việc nào")
	//		} else {
	//			for i, todos := range todo {
	//				fmt.Println("%d.%s \n", i+1, todos)
	//			}
	//		}
	//		fmt.Println("Danh sách việc làm", todo)
	//	case 2:
	//		var newtodo string
	//		fmt.Println("Thêm công việc mới")
	//		fmt.Scan(&newtodo)
	//		todo = append(todo, newtodo)
	//		fmt.Println("thêm công việc mới thành công :\n", todo)
	//	case 3:
	//		return
	//	}
	//}

	//	var n int
	//	fmt.Print("Nhập số lượng phần tử: ")
	//	fmt.Scan(&n)
	//
	//	arr := make([]int, n) // Tạo slice với n phần tử
	//	for i := 0; i < n; i++ {
	//		fmt.Printf("Nhập phần tử thứ %d: ", i+1)
	//		fmt.Scan(&arr[i])
	//	}
	//
	//	// Tính tổng
	//	sum := 0
	//	for _, value := range arr {
	//		sum += value
	//	}
	//	fmt.Printf("Tổng các phần tử: %d\n", sum)
	//
	//	var n int
	//	fmt.Println("Số lượng phần tử ")
	//	fmt.Scan(&n)
	//
	//	arr := make([]int, n)
	//	for i := 0; i < n; i++ {
	//		fmt.Println("Nhập Phần tử ", i+1)
	//		fmt.Scan(&arr[i])
	//	}
	//	max := arr[0]
	//	for _, v := range arr {
	//		if v > max {
	//			max = v
	//		}
	//	}
	//	fmt.Println("Số phần tử lớn nhất \n", max)

	//var n int
	//fmt.Println("Nhâp số phần tử")
	//fmt.Scan(&n)
	//
	//arr := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Println("Nhâp phần tử tiếp theo", i+n)
	//	fmt.Scan(&arr[i])
	//}
	//sochan := []int{}
	//for _, value := range arr {
	//	if value%2 == 0 {
	//		sochan = append(sochan, value)
	//	}
	//	fmt.Println(sochan)
	//}
	//sum := 0
	//for _, value := range sochan {
	//	sum += value
	//}
	//fmt.Println(sum)

	//	var num int
	//	fmt.Println("Nhập số nguyên:")
	//	fmt.Scan(&num)
	//
	//	if isParime(num) {
	//		fmt.Println(num, "là số nguyên tô")
	//	} else {
	//		fmt.Println(num, "Không phải là số nguyên tố")
	//	}
	//}

	//	var n int
	//	fmt.Println("Nhâp số Nguyên ")
	//	fmt.Scan(&n)
	//	if n < 1 {
	//		fmt.Println("không phải Số nguyên tố")
	//		return
	//	}
	//	fmt.Println("các số nguyên tố nhỏ hơn ", n)
	//
	//	for i := 1; i <= n; i++ {
	//		if isParime(i) {
	//			fmt.Println(i, "")
	//		}
	//
	//	}
	//
	//}
	//func isParime(n int) bool {
	//	if n < 2 {
	//		return false
	//	}
	//	for i := 2; i*i <= n; i++ {
	//		if n%i == 0 {
	//			return false
	//		}
	//	}
	//	return true

	// Kết nối cơ sở dữ liệu
	database.ConnecDB()
	r := router.SetupRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
