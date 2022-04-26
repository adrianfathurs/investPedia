package user

import "time"

type User struct {
	ID         int
	FullName   string
	Email      string
	Password   string
	Occupation string
	Avatar     string
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
