package entity

type LoanStatus struct {
	Id uint `gorm:"primaryKey"`
	Status string `gorm:"not null;type:varchar(10)"`
}
