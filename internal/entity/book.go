package entity

import "time"

type Book struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);index;not null"`
	Stock     uint   `gorm:"not null;default:0"`
	Author_id uint   `gorm:"not null"`
	Author    Author `gorm:"foreignKey:author_id;references:id"`
	Topic_id  uint   `gorm:"not null"`
	Topic     Topic  `gorm:"foreignKey:topic_id;references:id"`
	Cover_img string ``
	CreatedAt time.Time
	UpdatedAt time.Time
}
