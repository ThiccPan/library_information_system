package entity

type Role struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(10);not null"`
}
