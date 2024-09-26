package entity

type User struct {
	Id         string `gorm:"type:uuid;primaryKey"`
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	Name       string `gorm:"not null"`
	Role_id    uint   `gorm:"not null"`
	Role       Role   `gorm:"foreignKey:role_id;references:id"`
	ProfileImg string ``
}
