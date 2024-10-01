package entity

import "time"

type User struct {
	Id         uint   `gorm:"primaryKey"`
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	Name       string `gorm:"not null"`
	Role_id    uint   `gorm:"not null"`
	Role       Role   `gorm:"foreignKey:role_id;references:id"`
	Loans      []Loan
	ProfileImg string    ``
	CreatedAt  time.Time ``
	UpdatedAt  time.Time ``
}
