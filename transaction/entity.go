package transaction

import "time"

type transaction struct {
	ID                int
	Amount            int
	UserId            int
	TransactionStatus string
	Code              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
