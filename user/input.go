package user

type RegisterUserInput struct {
	FullName   string `json:"fullname" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Role       string
	Email      string `json:"email" binding:"required,email"`
}
