package entity

var (
	USER *Role = &Role{Id: 1, Name: "MEMBER"}
	ADMIN *Role = &Role{Id: 2, Name: "ADMIN"}
)

type Role struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(10);not null"`
}
