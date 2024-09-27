package entity

type Topic struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(50);not null"`
}
