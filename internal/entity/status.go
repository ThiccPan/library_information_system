package entity

type statusType string
var (
	BORROWED statusType = "BORROWED"
	RETURNED statusType = "RETURNED"
)

type LoanStatus struct {
	Id uint `gorm:"primaryKey"`
	Status statusType `gorm:"not null;type:varchar(10)"`
}
