package entity

type statusType string
type statusCode uint

const (
	BORROWED      statusType = "BORROWED"
	BORROWED_CODE statusCode = 1
	RETURNED      statusType = "RETURNED"
	RETURNED_CODE statusCode = 2
)

type LoanStatus struct {
	Id     uint       `gorm:"primaryKey"`
	Status statusType `gorm:"not null;type:varchar(10)"`
}
