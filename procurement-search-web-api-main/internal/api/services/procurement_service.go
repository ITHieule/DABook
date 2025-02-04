package services

import (
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/types"
)

type ProcurementService struct {
	*BaseService
}

var Procurement = &ProcurementService{}

func (s *ProcurementService) GetDataService(requestParams *request.ProcurementRequest) ([]types.YWCP, error) {
	var YWCPs []types.YWCP

	// Kết nối cơ sở dữ liệu
	db, err := database.RPConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn với tham số
	query := ` 
		SELECT *, DATEDIFF(DAY, Startdate, EndDate) AS DiffDay
		FROM (
			SELECT DDBH,
				   (SELECT TOP 1 USERDATE FROM KCRKS WHERE CGBH = ddzl.DDBH ORDER BY USERDATE) AS StartDate,
				   (SELECT TOP 1 INDATE FROM YWCP WHERE DDBH = ddzl.DDBH ORDER BY INDATE DESC) AS EndDate
			FROM (
				SELECT *,
					   (SELECT SUM(QTY) FROM YWCP WHERE DDBH = ddzl.DDBH) AS OKQTY
				FROM (
					SELECT DDBH, Pairs
					FROM ddzl
					WHERE LEFT(BUYNO, 6) = ? AND GSBH = 'VA12' AND pairs > 1
				) ddzl
			) ddzl
			WHERE pairs = OKQTY
		) YWCP1
		ORDER BY DiffDay
	`

	// Truyền tham số
	err = db.Raw(query, requestParams.BUYNO).Scan(&YWCPs).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}

	return YWCPs, nil
}

func (s *ProcurementService) GetAverageService(requestParams *request.ProcurementRequest) (float32, error) {

	// Kết nối cơ sở dữ liệu
	db, err := database.RPConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return 0, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn với tham số
	query := `
		SELECT CAST(AVG(CAST(s.DiffDay AS FLOAT)) AS DECIMAL(10, 4)) AS Average
		FROM (
			SELECT *, DATEDIFF(DAY, Startdate, EndDate) AS DiffDay
			FROM (
				SELECT DDBH,
					   (SELECT TOP 1 USERDATE FROM KCRKS WHERE CGBH = ddzl.DDBH ORDER BY USERDATE) AS StartDate,
					   (SELECT TOP 1 INDATE FROM YWCP WHERE DDBH = ddzl.DDBH ORDER BY INDATE DESC) AS EndDate
				FROM (
					SELECT *, (SELECT SUM(QTY) FROM YWCP WHERE DDBH = ddzl.DDBH) AS OKQTY
					FROM (
						SELECT DDBH, Pairs
						FROM ddzl
						WHERE LEFT(BUYNO, 6) = ? AND GSBH = 'VA12' AND Pairs > 1
					) ddzl
				) ddzl
				WHERE Pairs = OKQTY
			) YWCP1
		) AS s;
	`
	var average float32
	// Truyền tham số
	err = db.Raw(query, requestParams.BUYNO).Scan(&average).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return 0, err
	}

	return average, nil
}

func (s *ProcurementService) GetHelloService() ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	// Kết nối cơ sở dữ liệu
	db, err := database.RPConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn dữ liệu
	query := `
        SELECT top 100 * fROM WF_OUT
    `
	err = db.Raw(query).Scan(&results).Error
	if err != nil {
		fmt.Println("Query execution error:", err)
		return nil, err
	}

	return results, nil
}

func (s *ProcurementService) GetAccoutService(requestParams *request.Busers) (types.Busers, error) {
	var Busers types.Busers

	fmt.Println(requestParams)

	// Kết nối cơ sở dữ liệu
	db, err := database.RPConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return Busers, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn INSERT

	query := `
        INSERT INTO Busers (
            UserID, Username, PWD, Email, LastDateTime, Yn, PasswordChEnd,
            FromIP, DepID, EngName, Report, SupervisorID, Tyjh, Memo
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )
    `

	// Truyền tham số vào câu truy vấn
	err = db.Exec(query,
		requestParams.UserID,
		requestParams.Username,
		requestParams.Password,
		requestParams.Email,
		requestParams.LastDateTime,
		requestParams.Yn,
		requestParams.PasswordChEnd,
		requestParams.FromIP,
		requestParams.DepID,
		requestParams.EngName,
		requestParams.Report,
		requestParams.SupervisorID,
		requestParams.Tyjh,
		requestParams.Memo,
	).Error

	if err != nil {
		fmt.Println("Query execution error:", err)
	}

	return Busers, nil
}

func (s *ProcurementService) SearchDataByUserIDService(UserID string) ([]types.Busers, error) {
	var users []types.Busers

	// Kết nối cơ sở dữ liệu
	db, err := database.RPConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn với tham số
	query := `
		SELECT  UserID, Username, PWD, Email, LastDateTime, Yn, PasswordChEnd,
            FromIP, DepID, EngName, Report, SupervisorID, Tyjh, Memo
		FROM Busers
		WHERE UserID = ?
	`

	// Thực hiện truy vấn và ánh xạ kết quả vào cấu trúc dữ liệu
	err = db.Raw(query, UserID).Scan(&users).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}

	return users, nil
}
func (s *ProcurementService) DelteDataByUserIDService(UserID string) {

	// Kết nối cơ sở dữ liệu
	db, err := database.RPConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn với tham số
	query := `
		DELETE FROM Busers
		WHERE UserID = ?
	`

	// Thực hiện truy vấn và ánh xạ kết quả vào cấu trúc dữ liệu
	err = db.Raw(query, UserID).Scan(&UserID).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return
	}

	return
}
func (s *ProcurementService) UpdataDataByUserIDService(requestParams *request.Busers) ([]types.Busers, error) {
	var Busers []types.Busers

	// Kết nối cơ sở dữ liệu
	db, err := database.RPConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn với tham số
	query := `
		UPDATE Busers
		SET Username = ?, PWD = ?, Email = ?, LastDateTime = ?, Yn = ?, PasswordChEnd = ?, 
            FromIP = ?, DepID = ?, EngName = ?, Report = ?, SupervisorID = ?, Tyjh = ?, Memo = ?
		WHERE UserID = ?
	`

	// Thực hiện câu truy vấn
	err = db.Exec(query,
		requestParams.Username,
		requestParams.Password,
		requestParams.Email,
		requestParams.LastDateTime,
		requestParams.Yn,
		requestParams.PasswordChEnd,
		requestParams.FromIP,
		requestParams.DepID,
		requestParams.EngName,
		requestParams.Report,
		requestParams.SupervisorID,
		requestParams.Tyjh,
		requestParams.Memo,
		requestParams.UserID,
	).Error
	// Thực hiện truy vấn và ánh xạ kết quả vào cấu trúc dữ liệu

	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}

	return Busers, nil
}
