package user

type UserFormatter struct {
	ID         int    `json:"id"`
	FullName   string `json:"fullname"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

func FormatterRegisterResponse(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		FullName:   user.FullName,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}
	return formatter
}

func FormatterLoginResponse(user User, token string) UserFormatter {
	formatter := UserFormatter{
		FullName: user.FullName,
		Email:    user.Email,
		Token:    token,
	}
	return formatter
}
