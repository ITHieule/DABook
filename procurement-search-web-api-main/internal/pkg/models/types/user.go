package types

import "time"

type Busers struct {
	UserID        string    `json:"user_id" db:"USERID"`                // varchar(10)
	Username      string    `json:"username" db:"USERNAME"`             // varchar(30)
	Password      string    `json:"password" db:"PWD"`                  // varchar(10)
	Email         string    `json:"email" db:"EMAIL"`                   // varchar(40)
	LastDateTime  time.Time `json:"last_date_time" db:"LASTDATETIME"`   // smalldatetime
	Yn            string    `json:"yn" db:"YN"`                         // varchar(1)
	PasswordChEnd time.Time `json:"password_ch_end" db:"passwordchend"` // datetime
	FromIP        string    `json:"from_ip" db:"fromIP"`                // char(18)
	DepID         string    `json:"dep_id" db:"depid"`                  // varchar(4)
	EngName       string    `json:"eng_name" db:"Engname"`              // varchar(20)
	Report        string    `json:"report" db:"Report"`                 // varchar(100)
	SupervisorID  string    `json:"supervisor_id" db:"SupervisorID"`    // varchar(10)
	Tyjh          string    `json:"tyjh" db:"tyjh"`                     // varchar(1)
	Memo          string    `json:"memo" db:"Memo"`                     // varchar(20)
	//IsSSO         bool      `json:"is_sso" db:"IsSSO"`                  // bit
	//IsMFA         bool      `json:"is_mfa" db:"IsMFA"`                  // bit
	//MfaPwd        string    `json:"mfa_pwd" db:"MFA_PWD"`               // varchar(10)
	//MfaPwdExpire  time.Time `json:"mfa_pwd_expire" db:"MFA_PWD_Expire"` // smalldatetime
	//Resign_Date   time.Time `json:"resign_date" db:"Resign_Date"`       // smalldatetime
	//Role          string    `json:"role" db:"Role"`                     // varchar(1)
}
