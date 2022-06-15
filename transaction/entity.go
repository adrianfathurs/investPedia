package transaction

import (
	"investPedia/user"
	"time"
)

type Transaction struct {
	ID                int
	Amount            int
	UserId            int
	TransactionStatus string
	Code              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	User              user.User
}
