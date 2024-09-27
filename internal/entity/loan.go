package entity

import "time"

type Loan struct {
	Id            uint       `gorm:"primaryKey"`
	User_id       uint       `gorm:"not null"`
	User          User       `gorm:"foreignKey:user_id;references:id"`
	Book_id       uint       `gorm:"not null"`
	Book          Book       `gorm:"foreignKey:loan_status_id;references:id"`
	LoanStatus_id uint       `gorm:"not null"`
	LoanStatus    LoanStatus `gorm:"foreignKey:loan_status_id;references:id"`
	Deadline      time.Time
	CreatedAt     time.Time ``
	UpdatedAt     time.Time ``
}
