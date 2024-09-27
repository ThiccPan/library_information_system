package entity

import "time"

type Author struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);index;not null"`
	Bio       string ``
	CreatedAt time.Time
	UpdatedAt time.Time
}
